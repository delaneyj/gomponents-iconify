package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0V3a1 1 0 0 0-1-1ZM8 9h8v2a4 4 0 0 1-8 0V9Zm5 7.917A6.002 6.002 0 0 0 18 11V7H6v4a6.002 6.002 0 0 0 5 5.917V22a1 1 0 1 0 2 0v-5.083ZM14 3a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0V3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}