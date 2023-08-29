package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveTask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.964 7h-8v2h8V7ZM6 8.829v6.342L9.964 12L6 8.829ZM18.964 11h-8v2h8v-2Zm-8 4h8v2h-8v-2Z"/>`),
		g.Group(children),
	)
}