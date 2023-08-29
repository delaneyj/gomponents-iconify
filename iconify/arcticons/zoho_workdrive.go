package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZohoWorkdrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.024a2.176 2.176 0 0 1-2.18 2.171H6.672A2.176 2.176 0 0 1 4.5 37.017v-26.04a2.176 2.176 0 0 1 2.18-2.172v.001h9.902c2.257 0 6.225 4.13 8.188 4.238h16.908a1.822 1.822 0 0 1 1.822 1.821Z"/><circle cx="38.766" cy="22.935" r="3.084" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24.193" cy="28.564" r="3.084" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.48" cy="25.749" r="3.084" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.878 39.195v-2.873a3 3 0 0 1 1.923-2.8l5.07-1.95m.029 7.623v-5.611a3 3 0 0 1 1.922-2.8l5.07-1.95m.294 10.361V30.77a3 3 0 0 1 1.923-2.8l5.07-1.95"/>`),
		g.Group(children),
	)
}