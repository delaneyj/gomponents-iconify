package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7h6v2H7v2h3v2H7v4H5V7m8 0h6v2h-4v2h2a2 2 0 0 1 2 2v2c0 1.11-.89 2-2 2h-4v-2h4v-2h-4V7Z"/>`),
		g.Group(children),
	)
}