package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpDoubleArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 5H11l5 7l-5 7h4.5l5-7z"/><path fill="currentColor" d="M8.5 5H4l5 7l-5 7h4.5l5-7z"/>`),
		g.Group(children),
	)
}