package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
)

// comanda conectare la big query
// export GOOGLE_APPLICATION_CREDENTIALS=/Users/cristina/bqkey.json

func main() {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, "onedollarfootage")
	if err != nil {
		panic(err)
	}

	cins := client.Dataset("ejobsmap").Table("companies").Inserter()
	jins := client.Dataset("ejobsmap").Table("jobs").Inserter()

	var i int64
	for i = 309563; i <= 1000000; i++ {
		link := fmt.Sprintf("https://www.ejobs.ro/company/ejobs-group/%d", i)
		content, err := request(link)
		if err != nil {
			panic(err)
		}

		c, err := extract(i, *content)
		if err != nil {
			panic(err)
		}

		if c.Name != "" {
			log.Println(fmt.Sprintf("%s - %d jobs", c.Name, len(c.Jobs)))

			err = cins.Put(ctx, c)
			if err != nil {
				panic(err)
			}

			if len(c.Jobs) > 0 {
				err = jins.Put(ctx, c.Jobs)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
