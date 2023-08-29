package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraDocument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7v15H3v-2h16V7h-4.28c-.34.6-.98 1-1.72 1a2 2 0 0 1-2 2H8c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h3c1.1 0 2 .9 2 2c.74 0 1.38.41 1.72 1H19c1.11 0 2 .89 2 2M6 15h7l-2-4H8l-2 4Z"/>`),
		g.Group(children),
	)
}