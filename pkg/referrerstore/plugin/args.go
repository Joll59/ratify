package plugin

import (
	"fmt"
	"os"

	pluginCommon "github.com/deislabs/hora/pkg/common/plugin"
)

type ReferrerStorePluginArgs struct {
	Command          string
	Version          string
	SubjectReference string
	PluginArgs       [][2]string
	PluginArgsStr    string
}

var _ pluginCommon.PluginArgs = &ReferrerStorePluginArgs{}

func (args *ReferrerStorePluginArgs) AsEnv() []string {
	env := os.Environ()
	pluginArgsStr := args.PluginArgsStr
	if pluginArgsStr == "" {
		pluginArgsStr = pluginCommon.Stringify(args.PluginArgs)
	}

	// Duplicated values which come first will be overridden, so we must put the
	// custom values in the end to avoid being overridden by the process environments.
	// TODO replace the args
	env = append(env,
		"HORA_STORE_COMMAND="+args.Command,
		"HORA_STORE_SUBJECT="+args.SubjectReference,
		"HORA_STORE_ARGS="+pluginArgsStr,
		fmt.Sprintf("%s=%s", VersionEnvKey, args.Version),
	)
	return pluginCommon.DedupEnv(env)
}