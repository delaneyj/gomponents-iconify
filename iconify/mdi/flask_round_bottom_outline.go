package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskRoundBottomOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 15c0 1.66-1.34 3-3 3s-3-1.34-3-3v-1h2v1c0 .55.45 1 1 1s1-.45 1-1v-1h2v1m4 0c0 3.87-3.13 7-7 7s-7-3.13-7-7c0-2.79 1.64-5.2 4-6.33V5c0-.55.45-1 1-1h.5l-1-2h5l-1 2h.5c.55 0 1 .45 1 1v3.67c2.36 1.13 4 3.54 4 6.33m-8-9v4.1a5 5 0 1 0 2 0V6h-2Z"/>`),
		g.Group(children),
	)
}