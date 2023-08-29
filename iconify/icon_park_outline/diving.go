package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diving(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31.8 6H10.2C7.88 6 6 7.79 6 10v8h11c0-2 1.5-4 4-4s4 2 4 4h11v-8c0-2.21-1.88-4-4.2-4ZM16 24c0 1.491 1.25 6 5 6s5-4.509 5-6M42 6v32c0 4-3 6-6 6s-6-2-6-6v-2"/>`),
		g.Group(children),
	)
}