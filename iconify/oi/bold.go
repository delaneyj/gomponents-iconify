package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v1c.55 0 1 .45 1 1v4c0 .55-.45 1-1 1v1h5.5A2.5 2.5 0 0 0 8 5.5c0-1-.59-1.85-1.44-2.25A2 2 0 0 0 5 0H0zm3 1h1c.55 0 1 .45 1 1s-.45 1-1 1H3V1zm0 3h1.5C5.33 4 6 4.67 6 5.5S5.33 7 4.5 7H3V4z"/>`),
		g.Group(children),
	)
}