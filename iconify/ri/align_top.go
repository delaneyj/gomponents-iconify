package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18v2H3V3Zm5 8v10H6V11H3l4-4l4 4H8Zm10 0v10h-2V11h-3l4-4l4 4h-3Z"/>`),
		g.Group(children),
	)
}