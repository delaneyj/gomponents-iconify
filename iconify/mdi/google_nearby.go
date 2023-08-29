package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleNearby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.36 10.46l-7.82-7.82c-.85-.85-2.23-.85-3.08 0l-7.82 7.82c-.85.85-.85 2.23 0 3.08l7.82 7.82c.85.85 2.23.85 3.08 0l7.82-7.82c.85-.85.85-2.23 0-3.08M12 19l-7-7l7-7l7 7l-7 7m4.5-7L12 16.5L7.5 12L12 7.5l4.5 4.5Z"/>`),
		g.Group(children),
	)
}