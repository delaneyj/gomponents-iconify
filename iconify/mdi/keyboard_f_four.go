package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7h6v2H7v2h3v2H7v4H5V7m8 0h2v4h2V7h2v10h-2v-4h-4V7Z"/>`),
		g.Group(children),
	)
}