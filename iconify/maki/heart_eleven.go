package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10.06 4.76a27 27 0 0 1-4.2 5.36a.49.49 0 0 1-.71 0A27.003 27.003 0 0 1 .94 4.76C-.88 1.12 3.74-1.31 5.5 2.33c1.76-3.64 6.38-1.21 4.56 2.43z" fill="currentColor"/>`),
		g.Group(children),
	)
}