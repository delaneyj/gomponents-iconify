package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.334 24H2.668a2.668 2.668 0 0 1-2.667-2.667V2.666A2.669 2.669 0 0 1 2.668 0h18.667A2.668 2.668 0 0 1 24 2.666v18.667A2.668 2.668 0 0 1 21.333 24z"/>`),
		g.Group(children),
	)
}