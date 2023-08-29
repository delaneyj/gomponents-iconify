package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouchLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238 107.37V72a14 14 0 0 0-14-14H32a14 14 0 0 0-14 14v35.37A14 14 0 0 0 10 120v48a14 14 0 0 0 14 14h10v18a6 6 0 0 0 12 0v-18h164v18a6 6 0 0 0 12 0v-18h10a14 14 0 0 0 14-14v-48a14 14 0 0 0-8-12.63ZM226 72v34h-10a14 14 0 0 0-14 14v16a2 2 0 0 1-2 2h-66V70h90a2 2 0 0 1 2 2ZM32 70h90v68H56a2 2 0 0 1-2-2v-16a14 14 0 0 0-14-14H30V72a2 2 0 0 1 2-2Zm202 98a2 2 0 0 1-2 2H24a2 2 0 0 1-2-2v-48a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v16a14 14 0 0 0 14 14h144a14 14 0 0 0 14-14v-16a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}