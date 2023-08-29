package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3L5 7h3v7h2V7h3m3 10v-7h-2v7h-3l4 4l4-4h-3Z"/>`),
		g.Group(children),
	)
}