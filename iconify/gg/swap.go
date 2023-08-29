package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13v-1.5h-6v-2h6V8l3 2.5l-3 2.5Zm-8 4v-1.5h6v-2H8V12l-3 2.5L8 17Z"/>`),
		g.Group(children),
	)
}