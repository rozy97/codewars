package kata_test

import (
	. "github.com/rozy97/codewars/go/kata"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("There are 17 sheeps in total", func() {
		arr1 := []bool{
			true, true, true, false,
			true, true, true, true,
			true, false, true, false,
			true, false, false, true,
			true, true, true, true,
			false, false, true, true,
		}
		Expect(CountSheeps(arr1)).To(Equal(17))
	})
})
