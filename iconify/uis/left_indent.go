package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftIndent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zm0 4h10c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zm18 6H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1zM3 15h10c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zm18.8-5.3c-.4-.4-1-.5-1.4-.1l-2 1.7l-.1.1c-.4.4-.3 1.1.1 1.4l2 1.7c.2.1.4.2.6.2c.3 0 .6-.1.8-.4c.4-.4.3-1.1-.1-1.4l-1.1-.9l1.1-.9c.4-.4.4-1 .1-1.4z"/>`),
		g.Group(children),
	)
}