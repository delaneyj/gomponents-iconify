package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccumulationRain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 4a2.002 2.002 0 0 0-2 2v2h-2a3.438 3.438 0 0 1-3.051-2.316A1.007 1.007 0 0 0 20 5a.99.99 0 0 0-.947.679A3.44 3.44 0 0 1 16 8a3.438 3.438 0 0 1-3.051-2.316A1.007 1.007 0 0 0 12 5a.971.971 0 0 0-.947.679A3.44 3.44 0 0 1 8 8H6V6a2.002 2.002 0 0 0-2-2H2v2h2v20a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2V6h2V4ZM6 26v-4h4v-2H6v-4h4v-2H6v-4h2a4.932 4.932 0 0 0 4-1.987a5.02 5.02 0 0 0 8 0A4.932 4.932 0 0 0 24 10h2v16Z"/>`),
		g.Group(children),
	)
}