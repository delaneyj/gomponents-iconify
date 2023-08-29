package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7.2h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zm6 8H3c-.6 0-1 .4-1 1s.4 1 1 1h6c.6 0 1-.4 1-1s-.4-1-1-1zm9.5-5H3c-.6 0-1 .4-1 1s.4 1 1 1h15.5c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5h-2.8c.3-.4.4-.9 0-1.3s-1-.5-1.4-.1l-2 1.7l-.1.1c-.4.4-.3 1.1.1 1.4l2 1.7c.2.1.4.2.6.2c.3 0 .6-.1.8-.4c.3-.4.3-.9 0-1.3h2.8c1.9 0 3.5-1.6 3.5-3.5s-1.6-3.5-3.5-3.5z"/>`),
		g.Group(children),
	)
}