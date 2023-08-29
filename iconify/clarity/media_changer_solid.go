package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaChangerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path id="clarityMediaChangerSolid0" fill="currentColor" d="M30 4H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h1.88v1.57a1 1 0 0 0 2 0V32h16v1.57a1 1 0 0 0 2 0V32H30a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2ZM17 28H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm5 12h-2v-2h2Zm0-4h-2v-2h2Zm4 4h-2v-2h2Zm0-4h-2v-2h2Zm0-6h-6v-4h6Z"/>`),
		g.Group(children),
	)
}