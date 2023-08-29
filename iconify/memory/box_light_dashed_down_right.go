package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLightDashedDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M12 0v2h-2V0h2m0 4v3h-2V4h2m0 5v3H9v-2h1V9h2M.002 10H2v2H.002v-2M4 10h3v2H4v-2Z"/>`),
		g.Group(children),
	)
}