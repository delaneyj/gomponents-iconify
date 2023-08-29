package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 64v8a8 8 0 0 1-8 8H40a8 8 0 0 1-8-8v-8a8 8 0 0 1 8-8h176a8 8 0 0 1 8 8ZM40 120h128a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8H40a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8Zm176 16H40a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h176a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Zm-48 40H40a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h128a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}