package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 10v6h6v-6h4l-7-7l-7 7h4m3-4.2L14.2 8H13v6h-2V8H9.8L12 5.8M19 18H5v2h14v-2Z"/>`),
		g.Group(children),
	)
}