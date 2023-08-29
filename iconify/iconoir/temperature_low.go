package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 12a5 5 0 1 0 6 0m-6 0V3h6v9m0-9h2m-2 3h2m-2 3h2m5-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path d="M9 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0v-3"/></g>`),
		g.Group(children),
	)
}