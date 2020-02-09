package commit

import (
	"bytes"
	"errors"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/AlecAivazis/survey.v1/core"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var Cmd = &cobra.Command{
	Use:   "commit",
	Short: "通过交互式的问答，自动生成规范化的提交信息并提交变更",
	Run:   commitFunc,
}

var questions = []*survey.Question{
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "请选择本次提交的类型",
			Options: []string{
				"feat：添加了新的特性",
				"fix：修复Bug",
				"doc：添加或修改文档",
				"style：变更代码风格",
				"refactor：代码重构",
				"test：添加或变更测试",
				"build：构建工具或依赖变更",
				"script：辅助脚本变更",
			},
		},
		Validate: survey.Required,
		Transform: func(ans interface{}) interface{} {
			s := ans.(string)
			index := strings.Index(s, "：")
			return s[:index]
		},
	},
	{
		Name: "scope",
		Prompt: &survey.Input{
			Message: "请输入本次提交的作用域（建议填写组件或文件名）（可跳过）",
		},
	},
	{
		Name: "short",
		Prompt: &survey.Input{
			Message: "请输入本次提交的简要说明（不超过50字）",
		},
		Validate: func(ans interface{}) error {
			answer := strings.Trim(ans.(string), " ")
			if err := survey.Required(answer); nil != err {
				return errors.New("简要说明不能为空")
			}
			if err := survey.MaxLength(50)(answer); nil != err {
				return errors.New("简要说明的长度不能超过50个字")
			}

			return nil
		},
	},
	{
		Name: "long",
		Prompt: &survey.Multiline{
			Message: "请输入本次提交的详细说明（可跳过）",
		},
		Transform: func(ans interface{}) interface{} {
			return strings.Trim(ans.(string), "\n\n")
		},
	},
	{
		Name: "breaking",
		Prompt: &survey.Multiline{
			Message: "请输入本次提交的不兼容性变更（可跳过）",
		},
		Transform: func(ans interface{}) interface{} {
			return strings.Trim(ans.(string), "\n\n")
		},
	},
	{
		Name: "issue",
		Prompt: &survey.Multiline{
			Message: "请输入本次提交所修复或关闭的issue（建议每行输入一个相关issue，格式为'#123'）（可跳过）",
		},
		Transform: func(ans interface{}) interface{} {
			return strings.TrimRight(ans.(string), "\n\n")
		},
	},
}

func changeTemplate() {
	core.ErrorTemplate = errorOutputTemplate
	survey.SelectQuestionTemplate = selectQuestionTemplate
	survey.MultilineQuestionTemplate = multilineQuestionTemplate
}

type answer struct {
	Type     string
	Scope    string
	Short    string
	Long     string
	Breaking string
	Issue    string
}

func (a answer) String() string {
	var buffer bytes.Buffer
	err := template.Must(template.New("answer").Parse(answerFormatTemplate)).Execute(&buffer, a)
	if nil != err {
		panic(err)
	}
	return strings.TrimRight(buffer.String(), "\n")
}

func commitFunc(*cobra.Command, []string) {
	changeTemplate()

	var ans answer
	if err := survey.Ask(questions, &ans); nil != err {
		errorAction(err)
	}

	content := ans.String()

	var check bool
	if err := survey.AskOne(
		&survey.Confirm{
			Message: "\n##################################################\n" +
				content +
				"##################################################\n" +
				"请确认以上提交信息",
		}, &check, survey.Required); nil != err {
		errorAction(err)
	}

	command := exec.Command("git", "commit")
	if err := command.Run(); nil != err {
		errorAction(err)
	}
}

func errorAction(err error) {
	if err == terminal.InterruptErr {
		os.Exit(0)
	} else {
		panic(err)
	}
}
