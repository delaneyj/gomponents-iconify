package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.5 8.9l9 5.2c.2.1.3.1.5.1s.3 0 .5-.1l9-5.2c.2-.1.3-.2.4-.4c.2-.5.1-1.1-.4-1.4l-9-5.2c-.3-.2-.7-.2-1 0l-9 5.2c-.2.1-.3.2-.4.4c-.2.5-.1 1.1.4 1.4zm19 2.2l-.2-.1l-8.8 5.1c-.3.2-.7.2-1 0L2.7 11l-.2.1c-.5.3-.6.9-.4 1.4c.1.2.2.3.4.4l9 5.2c.3.2.7.2 1 0l9-5.2c.5-.3.6-.9.4-1.4c-.1-.2-.2-.3-.4-.4zm0 4l-.2-.1l-8.8 5.1c-.3.2-.7.2-1 0L2.7 15l-.2.1c-.5.3-.6.9-.4 1.4c.1.2.2.3.4.4l9 5.2c.3.2.7.2 1 0l9-5.2c.5-.3.6-.9.4-1.4c-.1-.2-.2-.3-.4-.4z"/>`),
		g.Group(children),
	)
}