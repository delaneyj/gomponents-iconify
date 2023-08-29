package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItalicA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m3 19l9.992-18H16l2 18h-3l-.4-5H8.5L6 19H3Zm7-8h4.5c-.051-.69-.483-6.429-.5-7c-.255.588-3.693 6.361-4 7Z"/>`),
		g.Group(children),
	)
}