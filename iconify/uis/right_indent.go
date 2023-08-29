package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightIndent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h18c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zm0 4h10c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zm18 6H3c-.6 0-1 .4-1 1s.4 1 1 1h18c.6 0 1-.4 1-1s-.4-1-1-1zM3 15h10c.6 0 1-.4 1-1s-.4-1-1-1H3c-.6 0-1 .4-1 1s.4 1 1 1zm15-1.3c0 .6.4 1 1 1c.2 0 .5-.1.6-.2l2-1.7l.1-.1c.4-.4.3-1.1-.1-1.4l-2-1.7c-.4-.4-1.1-.3-1.4.1c-.4.4-.3 1.1.1 1.4l1.1.9l-1.1.9c-.2.2-.3.5-.3.8z"/>`),
		g.Group(children),
	)
}