package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 1.75c-3.75 0-5.25 2-5.25 2v8.5s1.5 2 5.25 2s5.25-2 5.25-2v-8.5s-1.5-2-5.25-2z"/><path d="M2.75 8.25s1.5 2 5.25 2s5.25-2 5.25-2m-10.5-4s1.5 2 5.25 2s5.25-2 5.25-2"/></g>`),
		g.Group(children),
	)
}