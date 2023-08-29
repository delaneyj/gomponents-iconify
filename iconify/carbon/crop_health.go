package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropHealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 28h8v2H2zm15 1a1 1 0 0 1-.78-.375L12.52 24H2v-2h11a1 1 0 0 1 .78.375l3.147 3.932l5.241-7.862A1 1 0 0 1 23.8 18.4l2.7 3.6H30v2h-4a1 1 0 0 1-.8-.4l-2.152-2.87l-5.216 7.825a.999.999 0 0 1-.789.444zm-6-13v-5h1a4.005 4.005 0 0 0 4-4V4h-3a3.978 3.978 0 0 0-2.747 1.107A6.003 6.003 0 0 0 5 2H2v3a6.007 6.007 0 0 0 6 6h1v5H2v2h14v-2zm2-10h1v1a2.002 2.002 0 0 1-2 2h-1V8a2.002 2.002 0 0 1 2-2zM8 9a4.005 4.005 0 0 1-4-4V4h1a4.005 4.005 0 0 1 4 4v1z"/>`),
		g.Group(children),
	)
}