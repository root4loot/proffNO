# proffno

[![Build](https://github.com/root4loot/proffno/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/root4loot/proffno/actions/workflows/build.yml)

Get a list of subsidiaries for a Norwegian companies based on depth and ownership percentage level. Parses data from [proff.no](https://proff.no)

## Installation

```
go get github.com/root4loot/proffno
```

## Usage

```go
func main() {
	// Fetch subsidiaries of a company and its sub-subsidiaries owned by more than 50%
	corpData, err := proffno.FetchSubsidiaries("DnB Bank ASA", 2, 50)
	if err != nil {
		log.Fatalf("Failed to fetch subsidiaries: %v", err)
	}

	if corpData == nil {
		log.Println("No subsidiaries found")
		return
	}

	printSubsidiaries(corpData.Tree, 0)
}

// Recursively print subsidiaries
func printSubsidiaries(sub proffno.Subsidiary, level int) {
	fmt.Printf("%s%s (%.2f%%)\n", strings.Repeat("  ", level), sub.Name, sub.OwnedPercentage)
	for _, child := range sub.Sub {
		printSubsidiaries(child, level+1)
	}
}

```

### Output

```
DnB Bank ASA (0.00%)
  DnB Livsforsikring AS (100.00%)
    DnB Private Equity VI (is) AS (63.69%)
    DnB Private Equity IV (is) AS (58.95%)
    DnB Private Equity II (is) AS (55.27%)
    DnB Private Equity V (is) AS (72.55%)
    DnB Pe Direct II (is) AS (67.42%)
    DnB Næringseiendom AS (100.00%)
      DnB Ecp Invest AS (100.00%)
      DnB Ne Aif 1 AS (100.00%)
    DnB Eiendomsholding AS (100.00%)
      Lillebytunet AS (100.00%)
      DnB Handelsparker AS (100.00%)
      Bk9 Næring AS (100.00%)
      Nedre Skøyen Vei Newco AS (100.00%)
      Strandveien 50 AS (100.00%)
      Vitaminveien 1 AS (100.00%)
      Torgalmenningen 4 Hjemmel AS (100.00%)
      Starvhusgaten 2 B AS (100.00%)
      Torgalmenningen 14 AS (100.00%)
      Vitaminveien 1 Eiendom AS (100.00%)
      Strandgaten 4 Eiendom AS (100.00%)
      Rosenkrantzgaten 12 AS (100.00%)
      Markeveien 1 B AS (100.00%)
      Galleriet Kjøpesenter AS (100.00%)
      Sandslimarka 251 AS (100.00%)
      Grønvollkvartalet AS (100.00%)
      Vestnorsk Hotel AS (100.00%)
      Trondheim Torg AS (100.00%)
      Starvhusgaten 2 A AS (100.00%)
      Admiral Hotel AS (100.00%)
      Strandgaten 17 AS (100.00%)
      Beddingen 16 AS (100.00%)
      Stortingsgaten 22 AS (100.00%)
      Strandveien 50 AS (100.00%)
      Roald Amundsensgt 6 AS (100.00%)
      Brugata 19 AS (100.00%)
      Nordnorsk Hotell AS (100.00%)
      Fjordalléen 16 AS (100.00%)
      Trondheim Hotell AS (100.00%)
      DnB Eiendomsforvaltning AS (100.00%)
      Brugata 19 Hjemmel AS (100.00%)
      Roald Amundsensgt 6 Hjemmel AS (100.00%)
      Beddingen 16 Hjemmel AS (100.00%)
      Strandveien 50 Hjemmel AS (100.00%)
      Starvhusgaten 2 A Hjemmel AS (100.00%)
      Strandgaten 17 Hjemmel AS (100.00%)
      Rosenkrantzgaten 12 Hjemmel AS (100.00%)
      Strandgaten 4 Eiendom Hjemmel AS (100.00%)
      Stortingsgaten 22 Hjemmel AS (100.00%)
      Vestnorsk Hotel Hjemmel AS (100.00%)
      Starvhusgaten 2 B Hjemmel AS (100.00%)
      Hygea AS (100.00%)
      Markeveien 1 B Hjemmel AS (100.00%)
      Barcode 121 Hjemmel AS (100.00%)
      Pandox Tromsö AS (100.00%)
      Trondheim Hotell Hjemmel AS (100.00%)
      Trondheim Torg Hjemmel AS (100.00%)
    DnB Eiendomsinvest 2 AS (100.00%)
    DnB Kontor AS (100.00%)
      DnB Eiendomskomplementar AS (100.00%)
    DnB Kjøpesenter og Hotell AS (100.00%)
    DnB Liv Eiendom Sverige AS (100.00%)
  DnB Eiendomsutvikling AS (100.00%)
    Autolease AS (100.00%)
    Mosetertoppen Hafjell AS (100.00%)
    Skandinaviske Handelsparker AS (75.00%)
  Imove AS (100.00%)
  B&R Holding AS (60.00%)
    Uni Micro AS (66.67%)
      Uni Micro Web AS (77.50%)
      Profectus Invest AS (100.00%)
        Traveltext AS (91.37%)
  DnB Asset Management Holding AS (100.00%)
    DnB Asset Management AS (100.00%)
  DnB Invest Holding AS (100.00%)
  DnB Eiendom AS (100.00%)
  Godskipet 9 AS (100.00%)
  DnB Næringsmegling AS (100.00%)
  Godfjellet AS (100.00%)
    Yellow Holding AS (92.34%)
      Nille Holding AS (100.00%)
        Nille AS (100.00%)
  DnB Ventures AS (100.00%)
  Ocean Holding AS (100.00%)
    Godskipet 8 AS (100.00%)
  DnB Gjenstandsadministrasjon AS (100.00%)
  DnB Boligkreditt AS (100.00%)
```


