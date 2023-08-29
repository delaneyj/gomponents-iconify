package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m.146 3.854l.708-.708L5 7.293l3-3l6 6V5h1v7H8v-1h5.293L8 5.707l-3 3L.146 3.854Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}