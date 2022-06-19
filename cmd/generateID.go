package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/sony/sonyflake"
)

func GetID() string {

	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatal(err)
	}

	str := strconv.FormatUint(id, 16)

	fmt.Println(len(str))

	return str
}
