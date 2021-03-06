// Output-plug: outputstdout
// The plug's function take the event-data into the standard-output.
package outputstdout

import (
	"fmt"

	"github.com/wgliang/logcool/utils"
)

const (
	ModuleName = "stdout"
)

// Define outputstdout' config.
type OutputConfig struct {
	utils.OutputConfig
}

// Init outputstdout Handler.
func InitHandler(confraw *utils.ConfigRaw) (retconf utils.TypeOutputConfig, err error) {
	conf := OutputConfig{
		OutputConfig: utils.OutputConfig{
			CommonConfig: utils.CommonConfig{
				Type: ModuleName,
			},
		},
	}
	if err = utils.ReflectConfig(confraw, &conf); err != nil {
		return
	}

	retconf = &conf
	return
}

// Input's event,and this is the main function of output.
func (oc *OutputConfig) Event(event utils.LogEvent) (err error) {
	raw, err := event.MarshalIndent()
	if err != nil {
		return
	}

	fmt.Println(string(raw))
	return
}
