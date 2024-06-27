package prompt

import "github.com/manifoldco/promptui"

func RunPromptAuthor() string {
	pprompt := promptui.Prompt{
		Label: "Author",
	}
	result, err := pprompt.Run()
	if err != nil {
		panic(err)
	}
	return result
}

func RunPromptTitle() string {
	pprompt := promptui.Prompt{
		Label: "Title",
	}
	result, err := pprompt.Run()
	if err != nil {
		panic(err)
	}
	return result
}

func RunPromptStatus() string {
	prompt := promptui.Select{
		Label: "Status",
		Items: []string{"Read", "Reading", "To Read"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return result
}

func RunPromptNotFound() string {
	prompt := promptui.Select{
		Label: "Book not found",
		Items: []string{"search again", "add", "exit"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return result
}
