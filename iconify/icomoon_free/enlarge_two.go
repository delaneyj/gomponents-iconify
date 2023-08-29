package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnlargeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 0v6.5L13.5 4l-3 3L9 5.5l3-3L9.5 0zM7 10.5l-3 3L6.5 16H0V9.5L2.5 12l3-3z"/>`),
		g.Group(children),
	)
}