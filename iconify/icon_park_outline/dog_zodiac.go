package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DogZodiac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 44V19c0-5 3.6-9.4 14-15v9h7v6"/><path d="M16 25c4.013 1.78 11.354 5.124 13 15c.5 3 6 7 12 0c1.994-2.136 2.321-5.651-3.236-7.432"/><path d="M28 36c-3.333-.377-11 1-11 8"/></g>`),
		g.Group(children),
	)
}