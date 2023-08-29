package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceTextFormattingFilterOneFunnelFilterRoundOil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5.5H.5a6.51 6.51 0 0 0 5 6.33v6.67l3-2V6.83a6.51 6.51 0 0 0 5-6.33Z"/>`),
		g.Group(children),
	)
}