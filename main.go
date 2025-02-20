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
}

type CorporateData struct {
	Tree Subsidiary `json:"tree"`
}

func FetchSubsidiaries(orgName string, depth int, minOwnership float64) (*CorporateData, error) {
	if depth <= 0 {
		return nil, nil
	}

	orgName = formatCompanyName(orgName)
	orgNumber, err := retrieveCompanyInfo(orgName)
	if err != nil {
		return nil, err
	}

	corpData, err := retrieveCorporateStructure(orgNumber)
	if err != nil {
		return nil, err
	}

	filterSubsidiaries(&corpData.Tree, minOwnership)
	return corpData, nil
}

func retrieveCorporateStructure(orgNumber string) (*CorporateData, error) {
	resp, err := http.Get("https://proff.no/api/company/legal/" + orgNumber + "/corporateStructure")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var corpData CorporateData
	if err := json.NewDecoder(resp.Body).Decode(&corpData); err != nil {
		return nil, err
	}

	return &corpData, nil
}

func filterSubsidiaries(parent *Subsidiary, minOwnership float64) {
	filteredSubs := []Subsidiary{}
	for _, sub := range parent.Sub {
		if sub.OwnedPercentage > minOwnership {
			filterSubsidiaries(&sub, minOwnership)
			filteredSubs = append(filteredSubs, sub)
		}
	}
	parent.Sub = filteredSubs
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
