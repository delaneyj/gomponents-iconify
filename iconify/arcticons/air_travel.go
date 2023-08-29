package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirTravel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.353 33.245c2.734 0 4.903-.42 4.903-4.123V6.986h10.117v26.735c0 4.026-3.104 7.291-10.842 7.291"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.882 6.987h10.202v23.321c0 4.788 1.258 7.849 4.447 10.705c-8.972 0-14.65-1.218-14.65-8.478Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.632 17.867c-4.465 8.355 34.035 4.751 32.387-4.832"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.632 17.867c-24 14.258 44.652 10.286 32.387-4.832"/>`),
		g.Group(children),
	)
}