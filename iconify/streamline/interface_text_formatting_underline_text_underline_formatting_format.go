package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingUnderlineTextUnderlineFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.5.5v6.75a4.25 4.25 0 0 0 4.25 4.25h0A4.25 4.25 0 0 0 11 7.25V.5m-9.5 13h11"/>`),
		g.Group(children),
	)
}