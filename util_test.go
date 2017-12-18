package lg_test

import (
	"github.com/autopilothq/lg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExtractTrailingFields", func() {

	It("returns remaining args in the correct order", func() {
		args := []interface{}{1, 2, 3}
		_, remaining := lg.ExtractTrailingFields(args)
		Expect(len(remaining)).To(Equal(3))
		Expect(remaining[0]).To(Equal(1))
		Expect(remaining[1]).To(Equal(2))
		Expect(remaining[2]).To(Equal(3))
	})

})
