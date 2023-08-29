package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatIndentIncrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 7H4V5h16v2Zm0 4h-8V9h8v2Zm-8 4h8v-2h-8v2Zm-8 0l5-3l-5-3v6Zm0 2v2h16v-2H4Z"/>`),
		g.Group(children),
	)
}