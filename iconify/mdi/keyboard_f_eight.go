package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7h6v2H7v2h3v2H7v4H5V7m10 6v2h2v-2h-2m0-4v2h2V9h-2m0 8a2 2 0 0 1-2-2v-1.5c0-.83.67-1.5 1.5-1.5c-.83 0-1.5-.67-1.5-1.5V9c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1.5c0 .83-.67 1.5-1.5 1.5c.83 0 1.5.67 1.5 1.5V15c0 1.11-.89 2-2 2h-2Z"/>`),
		g.Group(children),
	)
}