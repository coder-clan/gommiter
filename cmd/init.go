package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gommiter/cmd/commit"
	"gommiter/cmd/hook"
	"gommiter/cmd/push"
	"os"
)

const version = "1.0.0"

var rootCmd = &cobra.Command{
	Long:    "一个简单的的命令行工具，用于规范化Git提交信息",
	Version: version,
}

func init() {
	rootCmd.AddCommand(commit.Cmd)
	rootCmd.AddCommand(push.PushCmd)
	//rootCmd.AddCommand(templateCmd)
	rootCmd.AddCommand(hook.HookCmd)
}

func Execute() {
	if err := rootCmd.Execute(); nil != err {
		fmt.Println(err)
		os.Exit(1)
	}
}
