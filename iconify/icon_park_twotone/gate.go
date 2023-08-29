package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGate0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M13 10v25m22-25v25"/><path d="M8 18h32"/><path stroke-linejoin="round" d="M24 10v8"/><path fill="#555" stroke-linejoin="round" d="M39 10H9L5 4s11.07 1 19 1s19-1 19-1l-4 6ZM10 35h6v9h-6zm22 0h6v9h-6z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGate0)"/>`),
		g.Group(children),
	)
}