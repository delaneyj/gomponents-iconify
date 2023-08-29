package feather

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudRain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 13v8m-8-8v8m4-6v8m8-6.42A5 5 0 0 0 18 7h-1.26A8 8 0 1 0 4 15.25"/>`),
		g.Group(children),
	)
}