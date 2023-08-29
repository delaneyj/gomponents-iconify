package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M8 16.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm7 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/><path d="M3 4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h1.05a2.5 2.5 0 0 1 4.9 0H10a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H3Zm11 3a1 1 0 0 0-1 1v6.05A2.5 2.5 0 0 1 15.95 16H17a1 1 0 0 0 1-1v-5a1 1 0 0 0-.293-.707l-2-2A1 1 0 0 0 15 7h-1Z"/></g>`),
		g.Group(children),
	)
}