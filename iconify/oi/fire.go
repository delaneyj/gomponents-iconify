package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M2 0c1 2-2 3-2 5s2 3 2 3c-.98-1.98 2-3 2-5S2 0 2 0zm3 3c1 2-2 3-2 5h3c.4 0 1-.5 1-2c0-2-2-3-2-3z"/>`),
		g.Group(children),
	)
}