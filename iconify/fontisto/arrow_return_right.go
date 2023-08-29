package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowReturnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.065 20.822H0V7.414h18.254V3.2L24 8.946l-5.746 5.746v-4.213H3.065v7.28h17.946v3.065z"/>`),
		g.Group(children),
	)
}