package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StairsUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 6h7v3h-4v4h-4v4h-4v4H3v-3h4v-4h4v-4h4V6m-4.83.66l-5.51 5.51l-1.83-1.83l5.51-5.51L6.5 3H12v5.5l-1.83-1.84Z"/>`),
		g.Group(children),
	)
}