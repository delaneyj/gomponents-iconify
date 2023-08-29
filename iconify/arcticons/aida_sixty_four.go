package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AidaSixtyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".937" d="M36.193 35.218V12.785L24.136 27.927h14.861"/><circle cx="16.568" cy="27.781" r="7.571" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".937"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".937" d="M23.298 15.443c-1.121-1.682-3.084-2.804-6.169-2.804h-.56a7.54 7.54 0 0 0-7.572 7.571v7.571"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}