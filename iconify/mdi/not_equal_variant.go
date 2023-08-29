package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotEqualVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.08 4.61l1.84.79L14.8 8H19v2h-5.05l-1.72 4H19v2h-7.62l-1.46 3.4l-1.84-.79L9.2 16H5v-2h5.06l1.71-4H5V8h7.63l1.45-3.39Z"/>`),
		g.Group(children),
	)
}