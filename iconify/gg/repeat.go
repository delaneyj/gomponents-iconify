package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.37 8l-4.5 2.598V9H6.89v4h-2V7h8.98V5.402L18.37 8Zm-8.24 9h8.98v-6h-2v4h-6.98v-1.598L5.63 16l4.5 2.598V17Z"/>`),
		g.Group(children),
	)
}