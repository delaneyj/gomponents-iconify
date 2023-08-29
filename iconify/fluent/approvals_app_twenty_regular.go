package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApprovalsAppTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.854 1.146a.5.5 0 1 0-.708.708L10.293 3H9.5a7.5 7.5 0 1 0 7.5 7.5a.5.5 0 0 0-1 0A6.5 6.5 0 1 1 9.5 4h.793L9.146 5.146a.5.5 0 1 0 .708.708l2-2a.5.5 0 0 0 0-.708l-2-2Zm3.493 5.994a.5.5 0 0 1 .013.707l-3.85 4a.5.5 0 0 1-.72 0l-1.65-1.715a.5.5 0 0 1 .72-.693l1.29 1.34l3.49-3.626a.5.5 0 0 1 .707-.013Z"/>`),
		g.Group(children),
	)
}