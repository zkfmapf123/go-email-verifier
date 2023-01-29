package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/zkfmapf123/email-verifier/src"
)

func main(){

	sc := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for sc.Scan() {
		m := src.CheckDomain(sc.Text())

		for k, v := range m {

			fmt.Println(k, "\t>\t", v)
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatal("[ERROR] : ", err)
	}

}
