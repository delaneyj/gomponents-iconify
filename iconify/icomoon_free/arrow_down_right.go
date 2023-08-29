package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m16 4.5l-4 4L3.5 0L0 3.5L8.5 12l-4 4H16V4.5z"/>`),
		g.Group(children),
	)
}