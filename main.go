package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-tools/go-steputils/stepconf"
	"github.com/bitrise-tools/go-steputils/tools"

	apiclient "github.com/peartherapeutics/bitrise-step-api-build-trigger/client"
	"github.com/peartherapeutics/bitrise-step-api-build-trigger/client/builds"
	"github.com/peartherapeutics/bitrise-step-api-build-trigger/models"

	httptransport "github.com/go-openapi/runtime/client"
)

var debugMode bool

// Config object mapping to bitrise step inputs
type Config struct {
	AppSlug  string          `env:"triggered_app_slug,required"`
	Token    stepconf.Secret `env:"bitrise_api_token,required"`
	Branch   string          `env:"triggered_branch,required"`
	Workflow string          `env:"triggered_workflow"`
	Message  string          `env:"trigger_message"`
	Verbose  bool            `env:"verbose,required"`
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

	log.Infof("Preparing API request...")

	/// API CLIENT
	// create the API client
	bitriseClient := apiclient.Default

	// Create auth provider
	apiKeyQueryAuth := httptransport.APIKeyAuth("Authorization", "header", string(cfg.Token))

	// Create and validate request parameters
	buildTriggerParams := builds.BuildTriggerParams{
		AppSlug: cfg.AppSlug,
		BuildParams: &models.V0BuildTriggerParams{
			BuildParams: &models.V0BuildTriggerParamsBuildParams{
				CommitMessage: cfg.Message,
				Branch:        cfg.Branch,
				WorkflowID:    cfg.Workflow,
			},
			HookInfo: &models.V0BuildTriggerParamsHookInfo{
				Type: "bitrise",
			},
		},
	}
	buildTriggerParams.SetTimeout(60 * time.Second) // Timeout defaults to 0 :(
	err := buildTriggerParams.BuildParams.Validate(strfmt.Default)
	if err != nil {
		failf("%s", err)
	}

	// Trigger build
	resp, err := bitriseClient.Builds.BuildTrigger(&buildTriggerParams, apiKeyQueryAuth)

	if err != nil {
		failf("%s", err)
	}
	log.Infof("\nSuccessfully started build: %#v\n", resp.Payload)
	// END

	// Output section
	exportVar("BITRISE_API_BUILD_TRIGGER_STEP_RESULT_BUILD_SLUG", resp.Payload.BuildSlug)
	exportVar("BITRISE_API_BUILD_TRIGGER_STEP_RESULT_BUILD_NUMBER", strconv.FormatInt(resp.Payload.BuildNumber, 10))
	exportVar("BITRISE_API_BUILD_TRIGGER_STEP_RESULT_BUILD_URL", resp.Payload.BuildURL)
	exportVar("BITRISE_API_BUILD_TRIGGER_STEP_RESULT_TRIGGERED_WORKFLOW", resp.Payload.TriggeredWorkflow)

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

func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	log.Warnf("For more details you can enable the debug logs by turning on the verbose step input.")
	os.Exit(1)
}

func exportVar(name, value string) {
	if err := tools.ExportEnvironmentWithEnvman(name, value); err != nil {
		failf("Failed to generate output - %s, %s", name, err)
	}
}
