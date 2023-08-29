package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.7 7.3c-.4-.4-1-.4-1.4 0L14 13.6L9.7 9.3C9.5 9.1 9.3 9 9 9c-.3 0-.5.1-.7.3l-6 6c-.2.2-.3.4-.3.7c0 .6.4 1 1 1c.3 0 .5-.1.7-.3L9 11.4l4.3 4.3c.1.1.2.2.3.2c.1.1.3.1.4.1c.2 0 .5-.1.6-.3h.1l7-7c.4-.4.4-1 0-1.4z"/>`),
		g.Group(children),
	)
}