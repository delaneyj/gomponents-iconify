package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kalgebra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.324 34.456h-8.657l4.333-6.5l-4.32-6.5h8.653"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.5 15.5h-9a2.006 2.006 0 0 1-2-2v-9h-18a2.006 2.006 0 0 0-2 2v35a2.006 2.006 0 0 0 2 2h27a2.006 2.006 0 0 0 2-2Zm-11-11l11 11"/>`),
		g.Group(children),
	)
}