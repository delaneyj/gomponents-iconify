package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MistSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19v-2h12v2H3Zm14 0v-2h4v2h-4ZM3 15v-2h4v2H3Zm6 0v-2h12v2H9Zm-6-4V9h13v2H3Zm15 0V9h3v2h-3ZM3 7V5h7v2H3Zm9 0V5h9v2h-9Z"/>`),
		g.Group(children),
	)
}