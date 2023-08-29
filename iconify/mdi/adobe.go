package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M14.58 3H22v16.67L14.58 3M9.42 3H2v16.67L9.42 3M12 9.17l4.67 10.5H13.5l-1.33-3.34H8.75L12 9.17z" fill="currentColor"/>`),
		g.Group(children),
	)
}