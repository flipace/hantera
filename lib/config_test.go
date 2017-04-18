package lib_test

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/flipace/hantera/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("lib.Config", func() {
	cwd, _ := filepath.Abs("/tmp")
	testFilePath := path.Join(cwd, "test.yml")

	It("can parse a valid yml file", func() {
		d1 := []byte("name: testconfig\ndependencies:\n  test-dependency:\n    description: demo description")
		err := ioutil.WriteFile(testFilePath, d1, 0644)

		Expect(err).NotTo(HaveOccurred())

		config := lib.GetProductConfig(testFilePath)

		Expect(config.Name).To(Equal("testconfig"))
		Expect(config.Dependencies).To(HaveKey("test-dependency"))
		Expect(config.Dependencies["test-dependency"].Description).To(Equal("demo description"))
	})

	It("panics when called with non-existent file", func() {
		testNonExistentFile := func() {
			lib.GetProductConfig("./nonexistent.yml")
		}

		Expect(testNonExistentFile).To(Panic())
	})

	It("panics when called on a malstructured file", func() {
		d1 := []byte("name: testconfig\ndependencies:\n  - this should break")
		err := ioutil.WriteFile(testFilePath, d1, 0644)

		Expect(err).NotTo(HaveOccurred())

		testMalformedFile := func() {
			lib.GetProductConfig(testFilePath)
		}

		Expect(testMalformedFile).To(Panic())
	})
})
