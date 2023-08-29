package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListNumberedRtlSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22v-1.5h2.5v-.75H18v-1.5h1.5v-.75H17V16h4v2.25L20 19l1 .75V22h-4Zm0-7v-3.75h2.5v-.75H17V9h4v3.75h-2.5v.75H21V15h-4Zm1.5-7V3.5H17V2h3v6h-1.5ZM3 19v-2h12v2H3Zm0-6v-2h12v2H3Zm0-6V5h12v2H3Z"/>`),
		g.Group(children),
	)
}