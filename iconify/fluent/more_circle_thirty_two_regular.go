package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreCircleThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm6 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm8-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm6 0c0 7.732-6.268 14-14 14S2 23.732 2 16S8.268 2 16 2s14 6.268 14 14Zm-2 0c0-6.627-5.373-12-12-12S4 9.373 4 16s5.373 12 12 12s12-5.373 12-12Z"/>`),
		g.Group(children),
	)
}