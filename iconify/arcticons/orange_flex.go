package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeFlex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-2 16.23l-6.987 9.258m6.987 0l-6.987-9.258"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.491 29.226a3.493 3.493 0 0 1-3.036 1.763h0a3.494 3.494 0 0 1-3.493-3.494v-2.27a3.494 3.494 0 0 1 3.494-3.495h0a3.494 3.494 0 0 1 3.493 3.494v1.136h-6.987M9.5 21.73h4.891m-2.82 9.258V19.457a2.446 2.446 0 0 1 2.446-2.446h.952a2.881 2.881 0 0 1 2.487 1.024v11.204a1.747 1.747 0 0 0 1.747 1.747h.524"/>`),
		g.Group(children),
	)
}