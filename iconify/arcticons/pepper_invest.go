package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PepperInvest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="1.5" height="1.5" x="38.212" y="38.108" fill="currentColor" rx=".75" ry=".75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.802 39.101h-3.25v-6.498h3.25m-3.25 3.249h2.112m2.814 2.518c.407.488.894.731 1.625.731h.975a1.63 1.63 0 0 0 1.625-1.624a1.63 1.63 0 0 0-1.625-1.625h-1.056a1.63 1.63 0 0 1-1.625-1.625a1.63 1.63 0 0 1 1.625-1.624h.975c.73 0 1.218.162 1.624.73m-10.965-.73L18.544 39.1l-2.112-6.498M10.231 39.1v-6.498l4.306 6.498v-6.498M35.51 39.1v-6.498m-2.112 0h4.225m-29.388 0V39.1"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}