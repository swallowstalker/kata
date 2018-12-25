package kata_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"

	"github.com/swallowstalker/positions-average/kata"
)

func assertFuzzy(act float64, exp float64) {
	var inrange bool
	var merr float64 = 1e-9
	e := math.Abs((act - exp) / exp)
	inrange = (e <= merr)
	if inrange == false {
		fmt.Printf("Expected should be near: %1.9e\n but got: %1.9e\n", exp, act)
	} else {
		//fmt.Println("GOOD")
	}
	Expect(inrange).To(Equal(true))
}
func dotest(s string, exp float64) {
	assertFuzzy(kata.PosAverage(s), exp)
}

var _ = Describe("PosAverage", func() {
	It("Basic tests", func() {
		dotest("466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096", 26.6666666667)
		dotest("444996, 699990, 666690, 096904, 600644, 640646, 606469, 409694, 666094, 606490", 29.2592592593)

	})
})
