package framework

import "flag"

type TestContextType struct {
	ReportPath string
}

var TestContext TestContextType

// RegisterCommonFlags registers flags common to all e2e test suites.
// The flag set can be flag.CommandLine (if desired) or a custom
// flag set that then gets passed to viperconfig.ViperizeFlags.
//
// The other Register*Flags methods below can be used to add more
// test-specific flags. However, those settings then get added
// regardless whether the test is actually in the test suite.
//
func RegisterCommonFlags(flags *flag.FlagSet) {
	flags.StringVar(&TestContext.ReportPath, "report-path", "", "The path of ginkgo report.")
}
