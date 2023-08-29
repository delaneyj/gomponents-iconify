package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePlus0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePlus1" fill="currentColor"><path id="fePlus2" d="M13 13v7a1 1 0 0 1-2 0v-7H4a1 1 0 0 1 0-2h7V4a1 1 0 0 1 2 0v7h7a1 1 0 0 1 0 2h-7Z"/></g></g>`),
		g.Group(children),
	)
}