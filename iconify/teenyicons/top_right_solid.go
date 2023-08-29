package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopRightSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 1h6v6h-1V2.707L1.854 13.854l-.708-.708L12.293 2H8V1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}