package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DogBowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 15l5.586-5.585A2 2 0 1 1 19 8a2 2 0 1 1-1.413 3.414L14 15"/><path d="M12 13L8.414 9.415A2 2 0 1 0 5 8a2 2 0 1 0 1.413 3.414L10 15m-7 5h18c-.175-1.671-.046-3.345-2-5H5c-1.333 1-2 2.667-2 5z"/></g>`),
		g.Group(children),
	)
}