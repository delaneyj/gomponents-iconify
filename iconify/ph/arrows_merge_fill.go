package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsMergeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 40v64a8 8 0 0 1-2.34 5.66L136 163.31V192h24a8 8 0 0 1 5.66 13.66l-32 32a8 8 0 0 1-11.32 0l-32-32A8 8 0 0 1 96 192h24v-28.69l-53.66-53.65A8 8 0 0 1 64 104V40a8 8 0 0 1 16 0v60.69l48 48l48-48V40a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}