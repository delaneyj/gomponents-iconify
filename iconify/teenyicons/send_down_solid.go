package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v1H1V1Zm7 2v9.293l3.146-3.147l.708.708L7.5 14.207L3.146 9.854l.708-.708L7 12.293V3h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}