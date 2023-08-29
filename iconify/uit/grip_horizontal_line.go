package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GripHorizontalLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 11h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1zm19 3h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}