package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DynamicFormOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 11V4h11v7H2Zm2-2h7V6H4v3ZM2 20v-7h13v7H2Zm2-2h9v-3H4v3Zm13 2v-9h-2V4h7l-2 5h2l-5 11ZM4.75 17.25h1.5v-1.5h-1.5v1.5Zm0-9h1.5v-1.5h-1.5v1.5ZM4 9V6v3Zm0 9v-3v3Z"/>`),
		g.Group(children),
	)
}