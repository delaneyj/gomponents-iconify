package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CongratulationHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 13.099L4.125 10.64v3.154s0 .548.355.633c1.114.264 2.279.516 3.52.516c1.27 0 2.485-.191 3.604-.543c.271-.058.271-.551.271-.551V10.64L8 13.099zM1.246 8.066h-.48C.355 9.63.052 14.603.052 14.912c0 .028.007.057.013.084h1.883a.395.395 0 0 0 .011-.084c-.001-.308-.305-5.281-.713-6.846z"/><path d="M8 1.062L0 6.043l1.33.999l5.555-1.258c.092-.529.528-.941 1.084-.941a1.125 1.125 0 0 1 0 2.25a1.07 1.07 0 0 1-.783-.352L2.374 7.83L8 11.702l8-5.659l-8-4.981z"/></g>`),
		g.Group(children),
	)
}