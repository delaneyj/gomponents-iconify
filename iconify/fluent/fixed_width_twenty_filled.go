package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FixedWidthTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 3a.5.5 0 0 1 .5.5V5h5.5V3.5a.5.5 0 0 1 1 0V5H16V3.5a.5.5 0 0 1 1 0v4a.5.5 0 0 1-1 0V6h-5.5v1.5a.5.5 0 0 1-1 0V6H4v1.5a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 .5-.5Zm6 6H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4.5V9Zm1 8H15a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-4.5v8Z"/>`),
		g.Group(children),
	)
}