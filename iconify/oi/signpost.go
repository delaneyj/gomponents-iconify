package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signpost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3 0v1H1L0 2l1 1h2v5h1V4h2l1-1l-1-1H4V0H3z"/>`),
		g.Group(children),
	)
}