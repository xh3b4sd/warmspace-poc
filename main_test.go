package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_Warmspace(t *testing.T) {
	testCases := []struct {
		des string
	}{
		{
			des: "case 000 description",
		},
	}

	for i := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var inp []byte
			{
				inp, err = os.ReadFile(prmFil(i))
				if err != nil {
					t.Fatal(err)
				}
			}

			var cur string
			{
				cur, err = call(string(inp))
				if err != nil {
					t.Fatal(err)
				}
			}

			var pat string
			{
				pat = gldFil(i)
			}

			{
				err := os.WriteFile(pat, []byte(cur), 0600)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func prmFil(i int) string {
	return fmt.Sprintf("prompt/case-%03d.txt", i)
}

func gldFil(i int) string {
	return fmt.Sprintf("golden/case-%03d.txt", i)
}
