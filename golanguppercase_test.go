package golanguppercase_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dremond71/golanguppercase"
)

var _ = Describe("golanguppercase", func() {
	Context("with affected strings", func() {
		It("hello", func() {
			Expect("HELLO").To(Equal(golanguppercase.Uppercase("hello")))
		})
		It("golang", func() {
			Expect("GOLANG").To(Equal(golanguppercase.Uppercase("golang")))
		})
		It("world", func() {
			Expect("WORLD").To(Equal(golanguppercase.Uppercase("world")))
		})
	})
	Context("with unaffected strings", func() {
		It("empty string", func() {
			Expect("").To(Equal(golanguppercase.Uppercase("")))
		})
		It("999", func() {
			Expect("999").To(Equal(golanguppercase.Uppercase("999")))
		})
		It("---", func() {
			Expect("---").To(Equal(golanguppercase.Uppercase("---")))
		})
		It("+++", func() {
			Expect("+++").To(Equal(golanguppercase.Uppercase("+++")))
		})
	})
})
