package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextColumnsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 184v8a8 8 0 0 1-8 8H40a8 8 0 0 1-8-8v-8a8 8 0 0 1 8-8h72a8 8 0 0 1 8 8Zm-8-88H40a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Zm0 40H40a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Zm0-80H40a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Zm104 40h-72a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Zm0 80h-72a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Zm0-120h-72a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Zm0 80h-72a8 8 0 0 0-8 8v8a8 8 0 0 0 8 8h72a8 8 0 0 0 8-8v-8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}