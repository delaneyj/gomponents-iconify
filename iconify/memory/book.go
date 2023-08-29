package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 2h1V1h14v1h1v18h-1v1H4v-1H3V2m8 7h-1V8H9v1H8v1H7V3H5v16h12V3h-5v7h-1V9Z"/>`),
		g.Group(children),
	)
}