package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cellar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 21h18v-9a9 9 0 1 0-18 0v9Zm0-4h18"/><path d="M9 17v-4h12m-8 0V9h7"/></g>`),
		g.Group(children),
	)
}