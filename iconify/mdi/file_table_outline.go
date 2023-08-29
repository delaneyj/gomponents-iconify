package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTableOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8l-6-6m4 18H6V4h7v5h5v11m-8-7H7v-2h3v2m4 0h-3v-2h3v2m-4 3H7v-2h3v2m4 0h-3v-2h3v2m-4 3H7v-2h3v2m4 0h-3v-2h3v2Z"/>`),
		g.Group(children),
	)
}