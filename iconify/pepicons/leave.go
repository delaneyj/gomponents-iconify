package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m15.667 8l2.083 2.5L15.667 8Zm0 5l2.083-2.5l-2.083 2.5Z" clip-rule="evenodd"/><path d="M16.5 10.5H10m-6-7h9m-9 14h9m0-14v4m0 6v4m-9-14v14"/></g>`),
		g.Group(children),
	)
}