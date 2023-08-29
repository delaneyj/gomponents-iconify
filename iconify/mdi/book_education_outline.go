package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookEducationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20h7v2H6c-1.11 0-2-.89-2-2V4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8.54l-1.5-.82l-.5.28V4h-5v8l-2.5-2.25L8 12V4H6v16m18-3l-5.5-3l-5.5 3l5.5 3l5.5-3m-9 2.09v2L18.5 23l3.5-1.91v-2L18.5 21L15 19.09Z"/>`),
		g.Group(children),
	)
}