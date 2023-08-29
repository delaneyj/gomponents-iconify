package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadSideCough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 16c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm-4 1c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zm3 3c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1zM16.2 2c-4-.1-7.2 2.9-7.3 6.9v.2l-1.8 3.8c-.2.5 0 1.1.5 1.3c.1.1.3.1.4.1h.9v1.8c0 1.1.9 1.9 1.9 1.9h.9v1.8c0 .6.4 1 1 1h8.7c.5-.2.8-.7.7-1.2l-.9-3L23 9.8v-.5c0-3.9-3-7.1-6.8-7.3z"/>`),
		g.Group(children),
	)
}