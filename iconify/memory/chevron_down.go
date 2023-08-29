package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M16 10h1V9h1V7h-2v1h-1v1h-1v1h-1v1h-1v1h-2v-1H9v-1H8V9H7V8H6V7H4v2h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h1v-1h1v-1h1v-1h1"/>`),
		g.Group(children),
	)
}