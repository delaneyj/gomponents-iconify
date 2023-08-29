package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m11.5 16l-4-4L16 3.5L12.5 0L4 8.5l-4-4V16h11.5z"/>`),
		g.Group(children),
	)
}