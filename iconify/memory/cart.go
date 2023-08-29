package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M19 14v2H6v-1H5v-4H4V8H3V3H1V1h4v3h16v4h-1v3h-1v1H7v2h12M5 7h1v3h12V7h1V6H5v1m2 10h2v1h1v2H9v1H7v-1H6v-2h1v-1m8 0h2v1h1v2h-1v1h-2v-1h-1v-2h1v-1Z"/>`),
		g.Group(children),
	)
}