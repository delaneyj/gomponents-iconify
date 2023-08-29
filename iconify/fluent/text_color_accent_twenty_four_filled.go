package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColorAccentTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 17a2 2 0 0 1 2-2h11.999a2 2 0 0 1 2 2v2.5a2 2 0 0 1-2 2H5.5a2 2 0 0 1-2-2V17Z"/>`),
		g.Group(children),
	)
}