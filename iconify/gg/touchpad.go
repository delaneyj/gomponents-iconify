package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Touchpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 21a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H4a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h16ZM4 5h16a1 1 0 0 1 1 1v8H3V6a1 1 0 0 1 1-1ZM3 16v2a1 1 0 0 0 1 1h7v-3H3Zm10 3v-3h8v2a1 1 0 0 1-1 1h-7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}