package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M144 56v144l-72-72Z" opacity=".2"/><path d="M224 120h-72V56a8 8 0 0 0-13.66-5.66l-72 72a8 8 0 0 0 0 11.32l72 72A8 8 0 0 0 152 200v-64h72a8 8 0 0 0 0-16Zm-88 60.69L83.31 128L136 75.31ZM48 40v176a8 8 0 0 1-16 0V40a8 8 0 0 1 16 0Z"/></g>`),
		g.Group(children),
	)
}