package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldLess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 3.586l4 4l4-4L17.414 5L12 10.414L6.586 5L8 3.586Zm4 10L17.414 19L16 20.414l-4-4l-4 4L6.586 19L12 13.586Z"/>`),
		g.Group(children),
	)
}