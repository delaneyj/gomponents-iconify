package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopLeftSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h6v1H2.707l11.147 11.146l-.708.708L2 2.707V7H1V1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}