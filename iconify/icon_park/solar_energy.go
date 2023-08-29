package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SolarEnergy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="40" height="24" x="4" y="8" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 32L30 8"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 32L18 8"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 20L6 20"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 41V32"/><path stroke-linecap="round" stroke-linejoin="round" d="M31 41H17"/></g>`),
		g.Group(children),
	)
}