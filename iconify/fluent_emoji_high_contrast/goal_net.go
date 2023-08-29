package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoalNet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.01 30a.54.54 0 0 1-.019 0H2.5a.5.5 0 0 1-.5-.5V15c0-1.824.781-3.126 1.919-3.939C5.009 10.281 6.357 10 7.5 10h17c1.143 0 2.49.282 3.581 1.061C29.218 11.874 30 13.176 30 15v14.5a.5.5 0 0 1-.5.5h-.491a.54.54 0 0 1-.018 0H28.5a.5.5 0 0 1-.5-.5v-.005L22.865 26.5H9.135L4 29.495v.005a.5.5 0 0 1-.5.5h-.49ZM4 15v3.338l4.5-2.625V12h-1c-.857 0-1.76.218-2.419.689C4.47 13.126 4 13.824 4 15Zm0 4.495v3.843l4.5-2.625V16.87L4 19.495Zm0 5v3.843l4.5-2.625V21.87L4 24.495Zm24 3.843v-3.843l-4.5-2.625v3.843l4.5 2.625Zm0-5v-3.843l-4.5-2.625v3.843l4.5 2.625Zm0-5V15c0-1.176-.468-1.874-1.081-2.311c-.66-.47-1.562-.689-2.419-.689h-1v3.713l4.5 2.625ZM22.5 12h-6v3.5h6V12Zm-7 0h-6v3.5h6V12Zm7 9.5h-6v4h6v-4Zm0-5h-6v4h6v-4Zm-7 0h-6v4h6v-4Zm0 5h-6v4h6v-4Z"/>`),
		g.Group(children),
	)
}