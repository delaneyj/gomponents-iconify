package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPlan0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M5 19h38v22a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V19Z"/><path stroke="#fff" stroke-linejoin="round" d="M5 10a2 2 0 0 1 2-2h34a2 2 0 0 1 2 2v9H5v-9Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m16 31l6 6l12-12"/><path stroke="#fff" stroke-linecap="round" d="M16 5v8m16-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPlan0)"/>`),
		g.Group(children),
	)
}