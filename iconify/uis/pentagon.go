package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.6 9.2l-9-6.5c-.4-.3-.8-.3-1.2 0l-9 6.5c-.3.2-.5.7-.4 1.1l3.4 10.6c.1.4.5.7 1 .7h11.1c.4 0 .8-.3 1-.7L22 10.3c.1-.4-.1-.9-.4-1.1z"/>`),
		g.Group(children),
	)
}