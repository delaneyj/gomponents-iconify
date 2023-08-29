package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrowCurvingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M53.214 10.786c-11.715-11.715-30.711-11.715-42.426 0c-11.717 11.715-11.717 30.711 0 42.426c11.715 11.717 30.711 11.717 42.426 0c11.715-11.715 11.715-30.711 0-42.426m-15.34 45.213v-5.166h-7.34c-9.504 0-17.533-8.883-17.533-19.398c0-10.717 8.344-19.436 18.6-19.436h9.207v8.176h-9.207c-5.943 0-10.777 5.053-10.777 11.26c0 6.084 4.447 11.223 9.711 11.223h7.34v-6.131l9.127 9.736l-9.128 9.736"/>`),
		g.Group(children),
	)
}