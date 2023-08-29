package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M37 23.2C37 32 24 44 24 44S11 32 11 23.2C11 15.91 16.82 10 24 10s13 5.91 13 13.2Z"/><path stroke-linejoin="round" stroke-miterlimit="2" d="M24 4v6m-6-5l4 5m8-5l-4 5M12 20h8m9 7h7m-20 7h6"/><path stroke-linejoin="round" d="M13.812 15A13.272 13.272 0 0 0 11 23.2c0 5.555 5.18 12.384 9 16.666"/><path d="M37 23.2c0 2.774-1.291 5.866-3.06 8.8"/></g>`),
		g.Group(children),
	)
}