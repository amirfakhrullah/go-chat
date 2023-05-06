package cli

import (
	"errors"
	"os"

	"github.com/manifoldco/promptui"
)

func GetApiKey(keyName string) (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
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

	os.Setenv(keyName, apiKey)
	return apiKey, nil
}

func GetInitialQuestion() (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("input is required")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Yes sir?",
		Validate: validate,
	}

	input, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return input, nil
}

func GetNextQuestion() (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New("input is required")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Next question sir? (Press `:q` to exit)",
		Validate: validate,
	}
	
	input, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return input, nil
}
