package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apron(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feApron0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feApron1" fill="currentColor" fill-rule="nonzero"><path id="feApron2" d="m9.694 8l-1 6H6v6h12v-6h-2.694l-1-6H9.694ZM10 3v3h4V3a1 1 0 0 1 2 0v3l1 6h1a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6a2 2 0 0 1 2-2h1l1-6V3a1 1 0 1 1 2 0ZM8 16h8v2H8v-2Z"/></g></g>`),
		g.Group(children),
	)
}