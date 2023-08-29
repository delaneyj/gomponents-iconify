package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineAirlineStops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8.7c-2.46 1.5-5.5 4.17-6 8.3h2v2H9v-2h2c-.5-4.5-4.37-8-9-8V7c4.39 0 8.22 2.55 10 6.3c1.38-2.97 3.86-5.03 5.96-6.31L14 7V5h7v7h-2V8.7z"/>`),
		g.Group(children),
	)
}