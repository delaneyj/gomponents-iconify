package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m8.867 8l2.083 2.5L8.867 8Zm0 5l2.083-2.5L8.867 13Z" clip-rule="evenodd"/><path d="M9.7 10.5H3.2m12.05 7h-9m9-14h-9m0 14v-4m0-6v-4m9.35 14v-14"/></g>`),
		g.Group(children),
	)
}