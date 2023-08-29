package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M2 0v1h1.63l-.06.13l-2 5l-.34.88H.01v1h5v-1H3.38l.06-.13l2-5L5.78 1H7V0H2z"/>`),
		g.Group(children),
	)
}