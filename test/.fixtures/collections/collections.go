package main

import (
	"fmt"
	"github.com/relops/cqlc/cqlc"
	"github.com/relops/cqlc/integration"
	"log"
	"os"
	"reflect"
)

func main() {

	s := integration.TestSession("127.0.0.1", "cqlc")
	cqlc.Truncate(s, COLLECTIONS)

	result := "FAILED"

	Input.Id = 12

	ctx := cqlc.NewContext()
	if err := ctx.Store(COLLECTIONS.Bind(Input)).Exec(s); err != nil {
		log.Fatalf("Could not store collections: %v", err)
		os.Exit(1)
	}

	var output Collections
	found, err := ctx.Select().From(COLLECTIONS).Where(COLLECTIONS.ID.Eq(Input.Id)).Into(COLLECTIONS.To(&output)).FetchOne(s)
	if err != nil {
		log.Fatalf("Could not store collections: %v", err)
		os.Exit(1)
	}

	if found {
		if reflect.DeepEqual(Input, output) {
			result = "PASSED"
		} else {
			result = fmt.Sprintf("Expected %+v but got %+v", Input, output)
		}
	}

	os.Stdout.WriteString(result)
}
