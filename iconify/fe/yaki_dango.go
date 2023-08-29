package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YakiDango(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feYakiDango0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feYakiDango1" fill="currentColor" fill-rule="nonzero"><path id="feYakiDango2" d="M13 17.83V21a1 1 0 0 1-2 0v-3.17a3.001 3.001 0 0 1-.659-5.33A2.997 2.997 0 0 1 9 10c0-1.043.533-1.963 1.341-2.5a3 3 0 1 1 3.318 0A2.997 2.997 0 0 1 15 10a2.997 2.997 0 0 1-1.341 2.5A3.001 3.001 0 0 1 13 17.83ZM12 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g></g>`),
		g.Group(children),
	)
}