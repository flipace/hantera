package develop_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHanteraDevelop(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("../../junit_commands_develop.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Hantera 'commands:develop' Suite", []Reporter{junitReporter})
}
