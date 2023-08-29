package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sandbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 12V9l6-6l3 3l-6 6H8zm10-7l-3-3l2-2l3 3l-2 2zM8 2h2v2H8zM4 2h2v2H4zM0 3a1 1 0 0 1 1-1h1v2H0V3zm0 3h2v2H0zm0 4h2v2H0zm0 4h2v2H0zm0 4h2v2H1a1 1 0 0 1-1-1v-1zm4 0h2v2H4zm4 0h2v2H8zm4 0h2v1a1 1 0 0 1-1 1h-1v-2zm0-4h2v2h-2z"/>`),
		g.Group(children),
	)
}