package develop_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	hantera "github.com/flipace/hantera/commands"
	develop "github.com/flipace/hantera/commands/develop"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/urfave/cli"
)

var _ = Describe("commands.develop.Setup", func() {
	app := hantera.GetApp("test", "test")

	targetPath, _ := filepath.Abs("/tmp/hantera")
	logFile := filepath.Join(targetPath, "hantera.log")

	var c *cli.Context

	BeforeEach(func() {
		c = cli.NewContext(app, nil, nil)

		os.MkdirAll(targetPath, os.ModePerm)

		d1 := []byte("name: testconfig\n" +
			"dependencies:\n" +
			"  themeit:\n" +
			"    repository: https://github.com/flipace/react-themeit\n")

		err := ioutil.WriteFile(filepath.Join(targetPath, "hantera.yml"), d1, 0644)

		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(targetPath)
	})

	It("throws when no hantera.yml can be found", func() {
		testFunc := func() {
			develop.Setup(c)
		}

		Expect(testFunc).To(Panic())
	})

	It("does not throw when config file passed via --config exists", func() {
		testFunc := func() {
			develop.RunSetup(
				filepath.Join(targetPath, "hantera.yml"),
				targetPath,
				"master", // branch
				true,     // no dependencies
				nil,      // context
			)
		}

		Expect(testFunc).To(Not(Panic()))

		os.RemoveAll(targetPath)
	})

	It("clones dependencies into the dependency-name folder beneath the target folder", func() {
		develop.RunSetup(
			filepath.Join(targetPath, "hantera.yml"),
			targetPath,
			"master", // branch
			true,     // no dependencies
			nil,      // context
		)

		Expect(filepath.Join(targetPath, "themeit")).To(BeADirectory())
		Expect(filepath.Join(targetPath, "themeit", "package.json")).To(BeARegularFile())
	})

	It("runs pre, override and post steps when defined", func() {
		// create a new yaml config
		d1 := []byte("name: testconfig\n" +
			"dependencies:\n" +
			"  themeit:\n" +
			"    repository: https://github.com/flipace/react-themeit\n" +
			"steps:\n" +
			"  setup:\n" +
			"    pre:\n      - echo \"pre run\" >> " + logFile + "\n" +
			"    override:\n      - echo \"override run\" >> " + logFile + "\n" +
			"    post:\n      - echo \"post run\" >> " + logFile + "\n")

		ioutil.WriteFile(filepath.Join(targetPath, "hantera.yml"), d1, 0644)

		develop.RunSetup(
			filepath.Join(targetPath, "hantera.yml"),
			targetPath,
			"master", // branch
			true,     // no dependencies
			nil,      // context
		)

		data, err := ioutil.ReadFile(logFile)

		Expect(err).To(Not(HaveOccurred()))

		Expect(string(data)).To(ContainSubstring("override run"), "override commands should have run")
		Expect(string(data)).To(ContainSubstring("pre run"), "pre commands should be run")
		Expect(string(data)).To(ContainSubstring("post run"), "post commands should be run")
	})
})
