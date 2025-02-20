package proffno

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Subsidiary struct {
	Name            string       `json:"name"`
	OwnedPercentage float64      `json:"ownedPercentage"`
	Sub             []Subsidiary `json:"sub"`
	Depth           int          `json:"-"`
}

type Result struct {
	InputQuery    string     `json:"inputQuery"`
	TargetCompany string     `json:"targetCompany"`
	Tree          Subsidiary `json:"tree"`
}

func (s *Subsidiary) OwnershipPercentage() float64 {
	return s.OwnedPercentage
}

func (s *Subsidiary) SubsidiaryName() string {
	return s.Name
}

func GetSubsidiaries(orgName string) (*Result, error) {
	orgName = formatCompanyName(orgName)
	orgNumber, err := retrieveCompanyInfo(orgName)
	if err != nil {
		return nil, err
	}

	result, err := retrieveCorporateStructure(orgNumber)
	if err != nil {
		return nil, err
	}

	result.InputQuery = orgName

	assignDepth(&result.Tree, 1) // root node is at depth 1
	return &Result{TargetCompany: orgName, Tree: result.Tree}, nil
}

// GetOwnedSubsidiaries returns a list of all companies that are owned (>50%) by the target company up to a specified depth
func (r *Result) GetOwnedSubsidiaries(maxDepth int) []string {
	var ownedSubsidiaries []string

	var collectOwned func(sub Subsidiary, parentOwned bool)
	collectOwned = func(sub Subsidiary, parentOwned bool) {
		if sub.Depth > maxDepth {
			return
		}

		isOwned := parentOwned || sub.OwnedPercentage > 50.0

		if isOwned {
			ownedSubsidiaries = append(ownedSubsidiaries, sub.Name)
		}

		for _, child := range sub.Sub {
			collectOwned(child, isOwned)
		}
	}

	collectOwned(r.Tree, false)

	return ownedSubsidiaries
}

func assignDepth(sub *Subsidiary, depth int) {
	sub.Depth = depth
	for i := range sub.Sub {
		assignDepth(&sub.Sub[i], depth+1)
	}
}

func retrieveCorporateStructure(orgNumber string) (*Result, error) {
	resp, err := http.Get("https://proff.no/api/company/legal/" + orgNumber + "/corporateStructure")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var Result Result
	if err := json.NewDecoder(resp.Body).Decode(&Result); err != nil {
		return nil, err
	}

	return &Result, nil
}

func retrieveCompanyInfo(query string) (string, error) {
	query = strings.ReplaceAll(query, " ", "+")
	buildID, err := fetchBuildID()
	if err != nil {
		return "", err
	}

	resp, err := http.Get("https://proff.no/_next/data/" + buildID + "/search.json?q=" + query)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var searchResponse struct {
		PageProps struct {
			HydrationData struct {
				SearchStore struct {
					Companies struct {
						Companies []struct {
							Orgnr string `json:"orgnr"`
						} `json:"companies"`
					} `json:"companies"`
				} `json:"searchStore"`
			} `json:"hydrationData"`
		} `json:"pageProps"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
		return "", err
	}

	if len(searchResponse.PageProps.HydrationData.SearchStore.Companies.Companies) == 0 {
		return "", nil
	}

	return searchResponse.PageProps.HydrationData.SearchStore.Companies.Companies[0].Orgnr, nil
}

func fetchBuildID() (string, error) {
	resp, err := http.Get("https://proff.no/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`"buildId":"(.*?)"`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", nil
	}
	return matches[1], nil
}

func formatCompanyName(name string) string {
	words := strings.Fields(name)
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		if (strings.ToLower(word) == "as" || strings.ToLower(word) == "asa") && i == len(words)-1 {
			words[i] = strings.ToUpper(word)
		}
	}
	return strings.Join(words, " ")
}
