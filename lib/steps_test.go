package lib_test

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/flipace/hantera/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("lib.Steps", func() {
	cwd, _ := filepath.Abs("/tmp")

	It("executes the passed slice of string commands", func() {
		commands := []string{
			"echo hello >/tmp/test.log && echo \"testing!\"",
			"echo hello2 >/tmp/test2.log && echo \"testing 2!\"",
		}

		lib.ExecuteStep(commands, cwd)

		dat1, err1 := ioutil.ReadFile(path.Join(cwd, "test.log"))
		dat2, err2 := ioutil.ReadFile(path.Join(cwd, "test2.log"))

		Expect(err1).To(Not(HaveOccurred()))
		Expect(err2).To(Not(HaveOccurred()))

		Expect(string(dat1)).To(Equal("hello\n"))
		Expect(string(dat2)).To(Equal("hello2\n"))
	})
})
