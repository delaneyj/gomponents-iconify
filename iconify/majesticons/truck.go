package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2h3a1 1 0 0 1 .707.293l4 4A1 1 0 0 1 22 12v5a1 1 0 0 1-1 1h-1.17a3 3 0 1 0-5.659 0H9.829a3 3 0 1 0-5.659 0H3a1 1 0 0 1-1-1V5zm7 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm8 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}