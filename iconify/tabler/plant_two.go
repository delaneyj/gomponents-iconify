package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlantTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 9a10 10 0 1 0 20 0"/><path d="M12 19A10 10 0 0 1 22 9M2 9a10 10 0 0 1 10 10"/><path d="M12 4a9.7 9.7 0 0 1 2.99 7.5m-5.98 0A9.7 9.7 0 0 1 12 4"/></g>`),
		g.Group(children),
	)
}