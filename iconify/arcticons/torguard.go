package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Torguard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.773 30.295c-5.312 2.11-10.27 6.225-12.036 12.205m14.813-8.932c-3.42 1.541-7.668 4.888-8.253 8.932m9.338 0H13.187m16.878-15.813c-8.823 1.31-14.612 8.464-16.878 15.813m11.997-21.786c-3.414 0-6.481 2.112-7.402 7.249V16.057a6.21 6.21 0 0 1 12.42 0v4.656h-5.018"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.635 42.5a6.765 6.765 0 0 0 1.92-13.252a4.048 4.048 0 0 0-3.004-5.685v.001v-7.506a10.559 10.559 0 0 0-21.117 0v11.265h-.246a7.589 7.589 0 0 0 0 15.177"/>`),
		g.Group(children),
	)
}