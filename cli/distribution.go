package cli

import "strings"

// BrandCommandHelp changes CLI-owned help, never response payloads or credentials.
func BrandCommandHelp(command *Command) {
	if CommandName == "aliyun" {
		return
	}
	render := func(s string) string { return strings.ReplaceAll(s, "aliyun ", CommandName+" ") }
	command.Usage = render(command.Usage)
	command.Sample = render(command.Sample)
	if command.Short != nil {
		for lang, text := range command.Short.GetData() {
			command.Short.GetData()[lang] = render(text)
		}
	}
	if command.Long != nil {
		for lang, text := range command.Long.GetData() {
			command.Long.GetData()[lang] = render(text)
		}
	}
	for _, child := range command.subCommands {
		BrandCommandHelp(child)
	}
}

// CommandText brands a CLI-owned command template, before interpolation.
func CommandText(text string) string { return strings.ReplaceAll(text, "aliyun ", CommandName+" ") }
