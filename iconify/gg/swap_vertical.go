package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 16h1.5v-6h2v6H17l-2.5 3l-2.5-3ZM8 8h1.5v6h2V8H13l-2.5-3L8 8Z"/>`),
		g.Group(children),
	)
}