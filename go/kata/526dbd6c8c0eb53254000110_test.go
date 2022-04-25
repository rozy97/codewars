// TODO: replace with your own tests (TDD). An example to get you started is included below.
// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo></http:>
// Gomega Matcher Library <http://onsi.github.io/gomega></http:>

package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(alphanumeric(".*?")).To(Equal(false))
		Expect(alphanumeric("a")).To(Equal(true))
		Expect(alphanumeric("Mazinkaiser")).To(Equal(true))
		Expect(alphanumeric("hello world_")).To(Equal(false))
		Expect(alphanumeric("PassW0rd")).To(Equal(true))
		Expect(alphanumeric("     ")).To(Equal(false))
		Expect(alphanumeric("")).To(Equal(false))
		Expect(alphanumeric("\n\t\n")).To(Equal(false))
		Expect(alphanumeric("ciao\n$$_")).To(Equal(false))
		Expect(alphanumeric("__ * __")).To(Equal(false))
		Expect(alphanumeric("&)))(((")).To(Equal(false))
		Expect(alphanumeric("43534h56jmTHHF3k")).To(Equal(true))
	})
})
