package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntertainmentNewsPaperNewspaperPeriodicalFoldContent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 4.5V11a1.25 1.25 0 0 1-1.25 1.25h0A1.25 1.25 0 0 1 11 11h0V2.25a.5.5 0 0 0-.5-.5H1a.5.5 0 0 0-.5.5v9a1 1 0 0 0 1 1h10.75"/><path d="M3.5 4.25H8v2.5H3.5zm0 5.5H8"/></g>`),
		g.Group(children),
	)
}