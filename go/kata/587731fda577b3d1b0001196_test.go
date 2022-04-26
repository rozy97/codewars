package kata_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rozy97/codewars/go/kata"
)

var _ = Describe("Basic tests", func() {
	t := [...][2]string{
		{"test case", "TestCase"},
		{"camel case method", "CamelCaseMethod"},
		{"say hello ", "SayHello"},
		{" camel case word", "CamelCaseWord"},
		{"", ""},
	}

	for _, v := range t {
		It(fmt.Sprintf("Testing input \"%s\"", v[0]), func() { Expect(CamelCase(v[0])).To(Equal(v[1])) })
	}
})
