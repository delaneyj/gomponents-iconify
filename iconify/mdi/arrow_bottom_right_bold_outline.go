package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomRightBoldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.88 15.54L4.22 9.88l5.66-5.66l5.66 5.66l4.24-4.24v14.14H5.64l4.24-4.24m7.78-4.95l-2.12 2.12l-5.66-5.66l-2.83 2.83l5.66 5.66l-2.12 2.12h7.07v-7.07Z"/>`),
		g.Group(children),
	)
}