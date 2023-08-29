package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pentagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.59 9.17l-9-6.54a1 1 0 0 0-1.18 0l-9 6.54a1 1 0 0 0-.36 1.12l3.44 10.58a1 1 0 0 0 1 .69h11.07a1 1 0 0 0 1-.69L22 10.29a1 1 0 0 0-.41-1.12Zm-4.75 10.39H7.16l-3-9.2L12 4.68l7.82 5.68Z"/>`),
		g.Group(children),
	)
}