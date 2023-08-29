package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 6h1V5a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v1h1v2.07l-5.71 9.9c-.19.3-.29.65-.29 1.03a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2c0-.38-.1-.73-.29-1.03L13 8.07V6M6 22a3 3 0 0 1-3-3c0-.6.18-1.16.5-1.63L9 7.81V7a1 1 0 0 1-1-1V5a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1v.81l5.5 9.56c.32.47.5 1.03.5 1.63a3 3 0 0 1-3 3H6Z"/>`),
		g.Group(children),
	)
}