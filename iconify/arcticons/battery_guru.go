package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryGuru(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.471a19.493 19.493 0 0 0 0 36.94"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.512 14.16A16.52 16.52 0 0 0 24 42.413M37.487 14.16A16.52 16.52 0 0 1 24 42.413"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.002 20.732c2.1 1.587 5.462 3.557 10.138-.539s8.711-4.35 15.253-2.02m-14.055 15.03l4.521-5.513h-4.367l4.521-5.513"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 42.412a19.493 19.493 0 0 0 0-36.94"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.56 14.072a16.484 16.484 0 0 1 2.928.089m-24.048-.089a16.484 16.484 0 0 0-2.928.089"/>`),
		g.Group(children),
	)
}