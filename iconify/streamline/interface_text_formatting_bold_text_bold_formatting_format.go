package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingBoldTextBoldFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 3.5a3 3 0 0 0-3-3H4a1 1 0 0 0-1 1v5h5a3 3 0 0 0 3-3Zm1 6.5a3.5 3.5 0 0 1-3.5 3.5H4a1 1 0 0 1-1-1v-6h5.5A3.5 3.5 0 0 1 12 10Z"/>`),
		g.Group(children),
	)
}