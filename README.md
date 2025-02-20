# proffno

[![Build](https://github.com/root4loot/proffno/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/root4loot/proffno/actions/workflows/build.yml)

Scrapes [proff.no](https://proff.no) to find subsidiaries for norwegian orgs.

## Installation

```
go get github.com/root4loot/proffno
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/root4loot/proffno"
)

func main() {
	results, err := proffno.GetSubsidiaries("DnB Bank ASA")
	if err != nil {
		log.Fatal(err)
	}
  
	processSubsidiaries(results.Tree)
}

func processSubsidiaries(sub proffno.Subsidiary) {
	indent := strings.Repeat("  ", sub.Depth)
	fmt.Printf("%s%d. %s (%.2f%%)\n", indent, sub.Depth+1, sub.Name, sub.OwnedPercentage)

	for _, child := range sub.Sub {
		processSubsidiaries(child)
	}
}
```

### Output
```
1. DnB Bank ASA (0.00%)
  2. DnB Livsforsikring AS (100.00%)
    3. DnB Private Equity VI (is) AS (63.69%)
    3. DnB Private Equity IV (is) AS (58.95%)
    3. DnB Private Equity II (is) AS (55.27%)
    3. DnB Private Equity V (is) AS (72.55%)
    3. DnB Pe Direct II (is) AS (67.42%)
    3. DnB Næringseiendom AS (100.00%)
      4. DnB Ecp Invest AS (100.00%)
      4. DnB Ne Aif 1 AS (100.00%)
    3. DnB Eiendomsholding AS (100.00%)
      4. Lillebytunet AS (100.00%)
      4. DnB Handelsparker AS (100.00%)
      4. Bk9 Næring AS (100.00%)
      4. Nedre Skøyen Vei Newco AS (100.00%)
      4. Strandveien 50 AS (100.00%)
      4. Vitaminveien 1 AS (100.00%)
      4. Torgalmenningen 4 Hjemmel AS (100.00%)
      4. Starvhusgaten 2 B AS (100.00%)
      4. Torgalmenningen 14 AS (100.00%)
      4. Vitaminveien 1 Eiendom AS (100.00%)
      4. Strandgaten 4 Eiendom AS (100.00%)
      4. Rosenkrantzgaten 12 AS (100.00%)
      4. Markeveien 1 B AS (100.00%)
      4. Galleriet Kjøpesenter AS (100.00%)
      4. Sandslimarka 251 AS (100.00%)
      4. Grønvollkvartalet AS (100.00%)
      4. Vestnorsk Hotel AS (100.00%)
      4. Trondheim Torg AS (100.00%)
      4. Starvhusgaten 2 A AS (100.00%)
      4. Admiral Hotel AS (100.00%)
      4. Strandgaten 17 AS (100.00%)
      4. Beddingen 16 AS (100.00%)
      4. Stortingsgaten 22 AS (100.00%)
      4. Strandveien 50 AS (100.00%)
      4. Roald Amundsensgt 6 AS (100.00%)
      4. Brugata 19 AS (100.00%)
      4. Nordnorsk Hotell AS (100.00%)
      4. Fjordalléen 16 AS (100.00%)
      4. Trondheim Hotell AS (100.00%)
      4. DnB Eiendomsforvaltning AS (100.00%)
      4. Brugata 19 Hjemmel AS (100.00%)
      4. Roald Amundsensgt 6 Hjemmel AS (100.00%)
      4. Beddingen 16 Hjemmel AS (100.00%)
      4. Strandveien 50 Hjemmel AS (100.00%)
      4. Starvhusgaten 2 A Hjemmel AS (100.00%)
      4. Strandgaten 17 Hjemmel AS (100.00%)
      4. Rosenkrantzgaten 12 Hjemmel AS (100.00%)
      4. Strandgaten 4 Eiendom Hjemmel AS (100.00%)
      4. Stortingsgaten 22 Hjemmel AS (100.00%)
      4. Vestnorsk Hotel Hjemmel AS (100.00%)
      4. Starvhusgaten 2 B Hjemmel AS (100.00%)
      4. Hygea AS (100.00%)
      4. Markeveien 1 B Hjemmel AS (100.00%)
      4. Barcode 121 Hjemmel AS (100.00%)
      4. Pandox Tromsö AS (100.00%)
      4. Trondheim Hotell Hjemmel AS (100.00%)
      4. Trondheim Torg Hjemmel AS (100.00%)
    3. DnB Eiendomsinvest 2 AS (100.00%)
    3. DnB Kontor AS (100.00%)
      4. DnB Eiendomskomplementar AS (100.00%)
    3. DnB Kjøpesenter og Hotell AS (100.00%)
    3. DnB Liv Eiendom Sverige AS (100.00%)
  2. DnB Eiendomsutvikling AS (100.00%)
    3. Autolease AS (100.00%)
    3. Mosetertoppen Hafjell AS (100.00%)
    3. Skandinaviske Handelsparker AS (75.00%)
  2. Imove AS (100.00%)
  2. B&R Holding AS (60.00%)
    3. Uni Micro AS (66.67%)
      4. Uni Micro Web AS (77.50%)
      4. Profectus Invest AS (100.00%)
        5. Traveltext AS (91.37%)
  2. DnB Asset Management Holding AS (100.00%)
    3. DnB Asset Management AS (100.00%)
  2. DnB Invest Holding AS (100.00%)
  2. DnB Eiendom AS (100.00%)
  2. Godskipet 9 AS (100.00%)
  2. DnB Næringsmegling AS (100.00%)
  2. Godfjellet AS (100.00%)
    3. Yellow Holding AS (92.34%)
      4. Nille Holding AS (100.00%)
        5. Nille AS (100.00%)
  2. DnB Ventures AS (100.00%)
  2. Ocean Holding AS (100.00%)
    3. Godskipet 8 AS (100.00%)
  2. DnB Gjenstandsadministrasjon AS (100.00%)
  2. DnB Boligkreditt AS (100.00%)
```


