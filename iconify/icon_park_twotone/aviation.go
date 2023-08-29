package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aviation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTAviation0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M32.5 30H36l8-8H11a3 3 0 0 0-1.8.6L4 26.5l6.277 3.177A3 3 0 0 0 11.63 30H13"/><path fill="#555" d="M27 31c0 1.657-3.582 3-8 3v-6c4.418 0 8 1.343 8 3Z"/><path fill="#555" stroke-linecap="round" d="m31 18l-9 4h15l3-10h-3l-6 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTAviation0)"/>`),
		g.Group(children),
	)
}