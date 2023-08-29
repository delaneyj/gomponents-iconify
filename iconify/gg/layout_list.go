package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7H7v2h2V7Zm-2 6v-2h2v2H7Zm0 2v2h2v-2H7Zm4 0v2h6v-2h-6Zm6-2v-2h-6v2h6Zm0-6v2h-6V7h6Z"/>`),
		g.Group(children),
	)
}