package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DynamicFormSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11V4h11v7H2Zm0 9v-7h13v7H2Zm15 0v-9h-2V4h7l-2 5h2l-5 11ZM4.75 17.25h1.5v-1.5h-1.5v1.5Zm0-9h1.5v-1.5h-1.5v1.5Z"/>`),
		g.Group(children),
	)
}