package main

import (
    "encoding/json"
	"fmt"
	"os"
    "path"
    "net/url"

    "github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
	"github.com/bitrise-tools/go-steputils/tools"
)

var debugMode bool

// Config ...
type Config struct {
	AppPath   string          `env:"app_path,required"`
	Token     stepconf.Secret `env:"appetize_token,required"`
	PublicKey string          `env:"public_key"`
	Verbose   bool            `env:"verbose,required"`
}

func main() {
    var cfg Config
	if err := stepconf.Parse(&cfg); err != nil {
		failf("Issue with input: %s", err)
	}

	debugMode = cfg.Verbose
	log.SetEnableDebugLog(debugMode)

	stepconf.Print(cfg)
	fmt.Println()

	log.Infof("Checking provided file's platform")

	fmt.Println("This is the value specified for the input 'example_step_input':", os.Getenv("example_step_input"))

    // Output section
	if err := tools.ExportEnvironmentWithEnvman("VERBOSE", "test exported value"); err != nil {
		failf("Failed to generate output - %s", "VERBOSE")
	}

	// You can find more usage examples on envman's GitHub page
	//  at: https://github.com/bitrise-io/envman

	//
	// --- Exit codes:
	// The exit code of your Step is very important. If you return
	//  with a 0 exit code `bitrise` will register your Step as "successful".
	// Any non zero exit code will be registered as "failed" by `bitrise`.
	os.Exit(0)
}

// -------------------------------------
// -- Private methods

func generateAppURL(publicKey string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "appetize.io",
		Path:   path.Join("app", publicKey),
	}
	return u.String()
}

func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	log.Warnf("For more details you can enable the debug logs by turning on the verbose step input.")
	os.Exit(1)
}

func logDebugPretty(v interface{}) {
	if !debugMode {
		return
	}

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	log.Debugf("Response: %+v\n", string(b))
}
