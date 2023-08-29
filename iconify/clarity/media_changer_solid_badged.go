package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaChangerSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityMediaChangerSolidBadged0" fill="currentColor"><path d="M30 13.5a7.49 7.49 0 0 1-4-1.16V14h-6v-4h3.66a7.49 7.49 0 0 1-1.16-4a7.37 7.37 0 0 1 .28-2H6a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h1.88v1.57a1 1 0 0 0 2 0V32h16v1.57a1 1 0 0 0 2 0V32H30a2 2 0 0 0 2-2V13.22a7.37 7.37 0 0 1-2 .28ZM17 28H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm0-4H8.81v-2H17Zm5 12h-2v-2h2Zm0-4h-2v-2h2Zm4 4h-2v-2h2Zm0-4h-2v-2h2Z"/><circle cx="30" cy="6" r="5"/></g>`),
		g.Group(children),
	)
}