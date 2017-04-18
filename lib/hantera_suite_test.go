package lib_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHantera(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("../junit_lib.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Hantera 'lib' Suite", []Reporter{junitReporter})
}
