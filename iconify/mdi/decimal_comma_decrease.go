package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecimalCommaDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 13v3H4l-1-3a1 1 0 0 1 2 0m10 3v-2l-3 3l3 3v-2h6v-2m-9-5a3 3 0 0 1-6 0V8a3 3 0 0 1 6 0m-2 0a1 1 0 0 0-2 0v3a1 1 0 0 0 2 0Z"/>`),
		g.Group(children),
	)
}