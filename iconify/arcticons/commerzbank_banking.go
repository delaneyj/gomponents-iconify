package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommerzbankBanking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.791 32.92H13.766c-3.058-.092-4.635-.947-6.196-3.502l-2.977-5.191L14.319 7.24l4.469-.045m8.054 4.133L36.855 28.67c1.45 2.695 1.498 4.488.065 7.117l-3.007 5.174l-19.573.069l-2.273-3.848m-.448-9.04L21.63 10.798c1.609-2.603 3.138-3.541 6.13-3.615l5.985.017l9.846 16.917l-2.194 3.892"/>`),
		g.Group(children),
	)
}