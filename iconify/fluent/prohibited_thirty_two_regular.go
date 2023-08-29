package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.837 23.749A11.952 11.952 0 0 1 4 16C4 9.373 9.373 4 16 4c2.954 0 5.658 1.067 7.749 2.837L6.837 23.749Zm1.414 1.414L25.163 8.251A11.953 11.953 0 0 1 28 16c0 6.627-5.373 12-12 12a11.953 11.953 0 0 1-7.749-2.837ZM2 16c0 7.732 6.268 14 14 14s14-6.268 14-14S23.732 2 16 2S2 8.268 2 16Z"/>`),
		g.Group(children),
	)
}