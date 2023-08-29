package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightIndentAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 17h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1zm0-4h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1zm-8-6h8c.6 0 1-.4 1-1s-.4-1-1-1h-8c-.6 0-1 .4-1 1s.4 1 1 1zm8 2h-8c-.6 0-1 .4-1 1s.4 1 1 1h8c.6 0 1-.4 1-1s-.4-1-1-1zM9 5c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1s1-.4 1-1V6c0-.6-.4-1-1-1zM3.6 9.6c-.4-.4-1-.3-1.4.1c-.4.4-.3 1.1.1 1.4l1.1.9l-1.1.9c-.2.2-.3.5-.3.8c0 .6.4 1 1 1c.2 0 .5-.1.6-.2l2-1.7l.1-.1c.4-.4.3-1.1-.1-1.4l-2-1.7z"/>`),
		g.Group(children),
	)
}