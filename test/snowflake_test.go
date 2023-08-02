package main

import (
	"fmt"
	snowflakeb "github.com/bwmarrin/snowflake"
	"go-simple-demo/pkg/snowflake"
	"testing"
	"time"
)

func Test_SnowflakeMine(t *testing.T) {
	worker0, _ := snowflake.NewWorker(0)
	worker00, _ := snowflake.NewWorker(0)
	worker1, _ := snowflake.NewWorker(1)
	fmt.Printf("%v,%v,%v\n", worker0.GetIDHex(), worker00.GetIDHex(), worker1.GetIDHex())
	fmt.Printf("%v,%v,%v\n", worker0.GetIDHex(), worker00.GetIDHex(), worker1.GetIDHex())
}
func Test_Snowflake_bwmarrin(t *testing.T) {
	st, _ := time.Parse("2006-01-02", "2023-08-01")
	snowflakeb.Epoch = st.UnixNano() / 1e6
	worker0, _ := snowflakeb.NewNode(0)
	worker00, _ := snowflakeb.NewNode(0)
	worker1, _ := snowflakeb.NewNode(1)

	// Generate a snowflake ID.
	fmt.Printf("%v,%v,%v\n", worker0.Generate(), worker00.Generate(), worker1.Generate())
	fmt.Printf("%v,%v,%v\n", worker0.Generate(), worker00.Generate(), worker1.Generate())
}
