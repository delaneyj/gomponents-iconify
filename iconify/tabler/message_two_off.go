package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 9h1m4 0h3m-8 4h5M8 4h10a3 3 0 0 1 3 3v8c0 .57-.16 1.104-.436 1.558M18 18h-3l-3 3l-3-3H6a3 3 0 0 1-3-3V7c0-1.084.575-2.034 1.437-2.561M3 3l18 18"/>`),
		g.Group(children),
	)
}