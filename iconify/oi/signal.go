package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M6 0v8h1V0H6zM4 1v7h1V1H4zM2 3v5h1V3H2zM0 5v3h1V5H0z"/>`),
		g.Group(children),
	)
}