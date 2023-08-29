package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M35.8 13H32L21 32L9.8 13H6"/><path d="M25.75 25.596c5.602 3.15 12.696 1.164 15.846-4.437c3.15-5.6 1.165-12.695-4.436-15.845c-5.601-3.15-12.695-1.164-15.846 4.437"/><path stroke-linejoin="round" d="M26 44H16m5 0V32m-9-16s2-2 5-2s5 3 8 3s5-1 5-1"/></g>`),
		g.Group(children),
	)
}