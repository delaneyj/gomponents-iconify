package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4.5 0C3.29 0 2.23.86 2 2C.9 2 0 2.9 0 4c0 .37.11.71.28 1H2.5l2-2l2 2h1.41c.06-.16.09-.32.09-.5c0-.65-.42-1.29-1-1.5v-.5A2.5 2.5 0 0 0 4.5 0zm0 4.5L2 7h2v.5a.5.5 0 1 0 1 0V7h2L4.5 4.5z"/>`),
		g.Group(children),
	)
}