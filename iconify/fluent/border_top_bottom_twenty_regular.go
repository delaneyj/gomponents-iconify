package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTopBottomTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.073 5.338a.509.509 0 0 0 .38.663a.509.509 0 0 0 .6-.462A2 2 0 0 1 6 4h8a2 2 0 0 1 1.947 1.54a.509.509 0 0 0 .6.462a.509.509 0 0 0 .38-.664A3.001 3.001 0 0 0 14 3H6a3 3 0 0 0-2.927 2.338Zm0 9.324a.509.509 0 0 1 .38-.663a.509.509 0 0 1 .6.462A2 2 0 0 0 6 16h8a2 2 0 0 0 1.947-1.54a.509.509 0 0 1 .6-.462a.509.509 0 0 1 .38.664A3.001 3.001 0 0 1 14 17H6a3 3 0 0 1-2.927-2.338ZM16.5 12a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 1 0v3a.5.5 0 0 1-.5.5ZM3 11.5a.5.5 0 0 0 1 0v-3a.5.5 0 0 0-1 0v3Z"/>`),
		g.Group(children),
	)
}