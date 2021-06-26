package readme

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"
)

var questions = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "Name of the project?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:      "description",
		Prompt:    &survey.Input{Message: "A Short Description of the project:"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name:   "supportString",
		Prompt: &survey.Input{Message: "How can users support you?"},
	},
	{
		Name: "motivation",
		Prompt: &survey.Multiline{
			Message: "Motivation for the project?",
		},
	},
	{
		Name: "contribute",
		Prompt: &survey.Multiline{
			Message: "How can one contribute?",
		},
	},
}

var fileExistsConfirmation = &survey.Confirm{
	Message: "File already exists do you want to overwrite it?",
}

func Setup() error {
	readmeDefinition := struct {
		Name          string
		Description   string
		SupportString string
		Motivation    string
		Contribute    string
	}{}
	err := survey.Ask(questions, &readmeDefinition)
	if err != nil {
		return err
	}

	var readmeString strings.Builder

	// Title
	readmeString.Write(
		[]byte("<h1 align=\"center\">" + readmeDefinition.Name + "</h1>\n"),
	)

	// Description
	readmeString.Write(
		[]byte("<p align=\"center\">" + readmeDefinition.Description + "</p>"),
	)

	// Support String
	readmeString.Write(
		[]byte("\n\n If you like any of my work, you can support me on: " + readmeDefinition.SupportString),
	)

	// Badges
	readmeString.Write(
		[]byte("\n\n[![](https://img.shields.io/badge/license-mit-black?style=for-the-badge)](LICENSE)\n\n"),
	)

	// Motivation
	readmeString.Write(
		[]byte("\n\n## Motivation \n\n"),
	)
	readmeString.Write(
		[]byte(readmeDefinition.Motivation),
	)

	// Features // TODO: need to add a questioning loop
	readmeString.Write(
		[]byte("\n\n## Features \n\n"),
	)
	for i := 1; i <= 3; i++ {
		readmeString.Write(
			[]byte("- Feature " + fmt.Sprint(i) + "\n"),
		)
	}

	// Contribute // TODO: allow a
	readmeString.Write(
		[]byte("\n\n## Contribute \n\n"),
	)
	readmeString.Write(
		[]byte(
			readmeDefinition.Contribute,
		),
	)

	// License // TODO: allow a
	readmeString.Write(
		[]byte("\n\n## License \n\n"),
	)
	readmeString.Write(
		[]byte(
			`[MIT](LICENSE) &copy; Reaper`,
		),
	)

	outputPath := "README.md"

	canOverwrite := true

	if _, err := os.Stat(outputPath); err == nil {
		survey.AskOne(fileExistsConfirmation, &canOverwrite, survey.WithValidator(survey.Required))
	}

	if canOverwrite {
		ioutil.WriteFile(outputPath, []byte(readmeString.String()), 0644)
	}

	return nil
}

func Run() {}
