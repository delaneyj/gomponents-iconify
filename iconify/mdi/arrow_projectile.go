package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowProjectile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 2l-2 5l-.97-.97L8 17.06V19l-3 3l-1-2l-2-1l3-3h1.94L17.97 4.97L17 4l5-2Z"/>`),
		g.Group(children),
	)
}