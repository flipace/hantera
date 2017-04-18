package lib_test

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/flipace/hantera/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("lib.Run", func() {
	It("executes the given command", func() {
		cwd, _ := filepath.Abs("/tmp")
		testFilePath := path.Join(cwd, "test.log")

		lib.Run(false, cwd, false, "echo \"it works\" > "+testFilePath)

		dat, err := ioutil.ReadFile(testFilePath)

		Expect(err).NotTo(HaveOccurred())
		Expect(string(dat)).To(Equal("it works\n"), "Expected \"it works\" to be written by Run")
	})
})
