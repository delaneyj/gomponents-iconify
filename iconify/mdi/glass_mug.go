package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassMug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 4v3h8V4h-8M8 2h13v1l-1 1v16l1 1v1H7v-1l1-1v-1.4l-3.8-1.77C3.5 16.5 3 15.82 3 15V8a2 2 0 0 1 2-2h3V4L7 3V2h1M5 15l3 1.39V8H5v7Z"/>`),
		g.Group(children),
	)
}