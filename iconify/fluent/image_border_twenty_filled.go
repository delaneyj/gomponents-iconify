package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageBorderTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 6h8v7.782l-2.802-2.788a1.7 1.7 0 0 0-2.396 0L6 13.782V6Zm5.5 1.5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-1.007 4.203L12.802 14H7.198l2.309-2.297a.7.7 0 0 1 .986 0ZM3 6a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Zm3-1a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H6Z"/>`),
		g.Group(children),
	)
}