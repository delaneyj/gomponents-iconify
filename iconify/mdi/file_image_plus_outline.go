package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImagePlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 19l5-5l1.88 1.88c-.55.91-.88 1.98-.88 3.12H7m3-8.5C10 9.67 9.33 9 8.5 9S7 9.67 7 10.5S7.67 12 8.5 12s1.5-.67 1.5-1.5m3.09 9.5H6V4h7v5h5v4.09c.33-.05.66-.09 1-.09c.34 0 .67.04 1 .09V8l-6-6H6a2 2 0 0 0-2 2v16c0 1.11.89 2 2 2h7.81c-.35-.61-.6-1.28-.72-2M18 15v3h-3v2h3v3h2v-3h3v-2h-3v-3h-2Z"/>`),
		g.Group(children),
	)
}