package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PuzzleCubeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 2v3h4V2H6ZM5 6H2v4h3V6Zm1 4V6h4v4H6Zm-1 1H2v.5A2.5 2.5 0 0 0 4.5 14H5v-3Zm1 3h4v-3H6v3Zm5 0v-3h3v.5a2.5 2.5 0 0 1-2.5 2.5H11Zm3-8v4h-3V6h3Zm0-1v-.5A2.5 2.5 0 0 0 11.5 2H11v3h3Z"/>`),
		g.Group(children),
	)
}