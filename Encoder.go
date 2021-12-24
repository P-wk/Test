/*
An Encoder according to 9.11.3.9A (5GS Update Type) in 3GPP TS 24501 (Version 16.9.0)
With coverage test
*/

package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

//The input structure
type Nas5GSUpdateType struct { 
	IEI             int
	Length          int
	EPS_PNB_CIoT    int
	FiveGS_PNB_CIoT int
	NG_RAN_RCU      int
	SMS_requested   int
}

//The encoder
func Encoder(ie Nas5GSUpdateType) *bytes.Buffer { 
	octet1 := strconv.FormatInt(int64(ie.IEI), 16)
	octet1 = Modify(octet1)
	octet2 := strconv.FormatInt(int64(ie.Length), 16)
	octet2 = Modify(octet2)
	temp := ie.SMS_requested + ie.NG_RAN_RCU*2 + ie.FiveGS_PNB_CIoT*4 + ie.EPS_PNB_CIoT*16
	octet3 := strconv.FormatInt(int64(temp), 16)
	octet3 = Modify(octet3)
	buffer := bytes.NewBufferString(octet1)
	buffer.WriteString(",")
	buffer.WriteString(octet2)
	buffer.WriteString(",")
	buffer.WriteString(octet3)
	return buffer
}

func Modify(s string) string { //A function used to modify output
	var s_mod string
	if len(s) == 1 {
		s_mod = "0x0" + s
	} else {
		s_mod = "0x" + s
	}
	return s_mod
}


//the coverage test. In this step, I use decimal random numbers to simulate all possible situations. Among them, the possible values of IEI and length are set at [0, 255].
func CodeCover(x int) { 
	var par []int
	var num int
	var number_s string
	var result *big.Int
	a := big.NewInt(int64(256))
	b := big.NewInt(int64(4))
	c := big.NewInt(int64(2))
	var randrange []*big.Int
	randrange = append(randrange, a)
	randrange = append(randrange, b)
	randrange = append(randrange, c)
	for i := 0; i <= x; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 2; k++ {
				var temp *big.Int
				temp = randrange[j]
				result, _ = rand.Int(rand.Reader, temp)
				number_s = result.String()
				num, _ = strconv.Atoi(number_s)
				par = append(par, num)
			}
		}
		ie := Nas5GSUpdateType{par[0], par[1], par[2], par[3], par[4], par[5]}
		bytestrom := Encoder(ie)
		par = []int{}
		fmt.Println(bytestrom)
	}

}

func main() {
	ie := Nas5GSUpdateType{1, 2, 0, 0, 1, 1}
	bytestrom := Encoder(ie)
	fmt.Println("Output:", bytestrom)
	fmt.Println("\n Test for coverage \n")
	tsetNum := 500 //number of samples used to test coverage
	CodeCover(tsetNum)
}
