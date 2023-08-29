package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18v2H3V3Zm9 2.586l-.707.707l-4 4l-.707.707L8 12.414l.707-.707L11 9.414V21h2V9.414l2.293 2.293l.707.707L17.414 11l-.707-.707l-4-4L12 5.586Z"/>`),
		g.Group(children),
	)
}