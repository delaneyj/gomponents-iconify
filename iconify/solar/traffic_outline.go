package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrafficOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M16.114 21.932A10.75 10.75 0 1 1 12 1.25a.75.75 0 0 1 0 1.5A9.25 9.25 0 1 0 21.25 12a.75.75 0 0 1 1.5 0a10.75 10.75 0 0 1-6.636 9.932Z"/><path d="M14.687 1.589a.75.75 0 1 0-.374 1.452a9.267 9.267 0 0 1 6.646 6.646a.75.75 0 0 0 1.452-.374a10.768 10.768 0 0 0-7.724-7.724Z"/></g>`),
		g.Group(children),
	)
}