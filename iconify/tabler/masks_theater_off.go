package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MasksTheaterOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 9h6.808a2 2 0 0 1 1.992 2.183l-.554 6.041m-1.286 2.718A3.99 3.99 0 0 1 17.25 21h-1.5a4 4 0 0 1-3.983-3.635l-.567-6.182M18 13h.01"/><path d="M15 16.5c.657.438 1.313.588 1.97.451m-8.338-.969A4.05 4.05 0 0 1 8.25 16h-1.5a4 4 0 0 1-3.983-3.635L2.2 6.183a2 2 0 0 1 .514-1.531A1.99 1.99 0 0 1 4 4m4 0h2.808a2 2 0 0 1 2 2M6 8h.01"/><path d="M6 12c.764-.51 1.528-.63 2.291-.36M3 3l18 18"/></g>`),
		g.Group(children),
	)
}