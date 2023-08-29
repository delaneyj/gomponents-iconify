package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingTextStrikeThroughTextStrikeThroughFormattingFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 .5V5m-2.5 8.5h5M1.5 3V1.5a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1V3m-11 4.5h11M7 7.5v6"/>`),
		g.Group(children),
	)
}