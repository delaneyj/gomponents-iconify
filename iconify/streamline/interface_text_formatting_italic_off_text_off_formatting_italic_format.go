package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingItalicOffTextOffFormattingItalicFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.24 6.24L4 13.5m-2.5 0h5M5 .5h7.5a1 1 0 0 1 1 1V3M.5.5l13 13M8 .5L7.23 3"/>`),
		g.Group(children),
	)
}