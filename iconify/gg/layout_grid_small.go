package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutGridSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 7h2v2H7V7Zm4 0h2v2h-2V7Zm6 0h-2v2h2V7ZM7 11h2v2H7v-2Zm6 0h-2v2h2v-2Zm2 0h2v2h-2v-2Zm-6 4H7v2h2v-2Zm2 0h2v2h-2v-2Zm6 0h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}