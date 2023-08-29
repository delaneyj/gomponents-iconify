package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1.764l9.946 19.892L12 18.548l-9.946 3.108L12 1.764Zm0 4.472L5.946 18.344L12 16.452l6.054 1.892L12 6.236Z"/>`),
		g.Group(children),
	)
}