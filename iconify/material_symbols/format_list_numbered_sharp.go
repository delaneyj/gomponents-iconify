package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListNumberedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-1.5h2.5v-.75H4v-1.5h1.5v-.75H3V16h4v2.25L6 19l1 .75V22H3Zm0-7v-3.75h2.5v-.75H3V9h4v3.75H4.5v.75H7V15H3Zm1.5-7V3.5H3V2h3v6H4.5ZM9 19v-2h12v2H9Zm0-6v-2h12v2H9Zm0-6V5h12v2H9Z"/>`),
		g.Group(children),
	)
}