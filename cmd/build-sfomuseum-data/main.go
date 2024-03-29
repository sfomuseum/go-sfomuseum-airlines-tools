package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/sfomuseum/go-sfomuseum-airlines-tools"
	"github.com/sfomuseum/go-sfomuseum-airlines-tools/template"
	"log"
	"os"
)

func main() {

	iterator_uri := flag.String("iterator-uri", "repo://", "...")
	iterator_source := flag.String("iterator-source", "/usr/local/data/sfomuseum-data-enterprise", "...")

	flag.Parse()

	ctx := context.Background()

	lookup, err := tools.CompileAirlinesData(ctx, *iterator_uri, *iterator_source)

	if err != nil {
		log.Fatalf("Failed to compile data, %v", err)
	}

	enc, err := json.Marshal(lookup)

	if err != nil {
		log.Fatal(err)
	}

	vars := template.AirlineDataVars{
		Package: "sfomuseum",
		Data:    string(enc),
	}

	err = template.RenderAirlineData(os.Stdout, &vars)

	if err != nil {
		log.Fatal(err)
	}
}
