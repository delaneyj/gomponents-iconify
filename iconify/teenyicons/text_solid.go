package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h13v4h-1V2H8v11h3v1H3.5v-1H7V2H2v3H1V1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}