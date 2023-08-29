package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaRecord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 1C2.34 1 1 2.34 1 4s1.34 3 3 3s3-1.34 3-3s-1.34-3-3-3z"/>`),
		g.Group(children),
	)
}