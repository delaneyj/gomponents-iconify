package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignSpaceFitVerticalTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1a2 2 0 0 0-2 2v4.5a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2H6ZM5 3a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v4.5a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V3Zm1 7.5a2 2 0 0 0-2 2V17a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-4.5a2 2 0 0 0-2-2H6Zm-1 2a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1V17a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-4.5Z"/>`),
		g.Group(children),
	)
}