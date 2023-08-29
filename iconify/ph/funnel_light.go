package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228.77 50.34A13.8 13.8 0 0 0 216 42H40a14 14 0 0 0-10.33 23.42l.06.07l67.73 72.31a2 2 0 0 1 .54 1.37V216a14 14 0 0 0 21.77 11.64l32-21.33a14 14 0 0 0 6.23-11.65v-55.49a2 2 0 0 1 .54-1.37l67.79-72.38a13.82 13.82 0 0 0 2.44-15.08Zm-11.26 6.94l-67.73 72.32a13.93 13.93 0 0 0-3.78 9.57v55.49a2 2 0 0 1-.89 1.67l-32 21.33A2 2 0 0 1 110 216v-76.83a14 14 0 0 0-3.78-9.58L38.53 57.32A2 2 0 0 1 40 54h176a1.9 1.9 0 0 1 1.83 1.19a1.86 1.86 0 0 1-.32 2.09Z"/>`),
		g.Group(children),
	)
}