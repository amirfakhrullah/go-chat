package cli

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func GetApiKey(keyName string) (string, error) {
	validate := func(input string) error {
		if len(input) == 5 {
			return errors.New("open ai api key is required")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Open AI API key is not set yet. Insert the key to get started",
		Validate: validate,
	}

	apiKey, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return apiKey, nil
}

func GetQuestion(firstQuestion bool) (string, error) {
	validate := func(input string) error {
		if (input != ":q!" && len(input) < 6) {
			return errors.New("input must be at least 6 characters")
		}
		return nil
	}

	label := "Next question? (Press `:q!` to exit)"
	if (firstQuestion) {
		label = "Your question?"
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	input, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return input, nil
}
