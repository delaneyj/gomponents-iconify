package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tradingview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 14.453v7.82h7.365v11.274h7.82V14.453H4.5zm21.702 19.094l8.124-19.094H43.5l-8.124 19.094h-9.174z"/><circle cx="25.841" cy="18.363" r="3.91" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}