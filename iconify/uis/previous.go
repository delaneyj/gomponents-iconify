package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.7 15.3L13.4 12l3.3-3.3c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0l-4 4c-.4.4-.4 1 0 1.4l4 4c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.4-.4.4-1 0-1.4zM8 7c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1s1-.4 1-1V8c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}