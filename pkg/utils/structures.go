package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrintByteSlice(structure interface{}) (string, error) {
	bytes, err := json.MarshalIndent(structure, "", "\t")
	if err != nil {
		e := ReturnError("PrettyPrintByteSlice:11", err)
		return "", e
	}

	structureAsString := string(bytes)
	fmt.Print(structureAsString)

	return structureAsString, nil
}
