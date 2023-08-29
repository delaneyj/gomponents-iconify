package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckRectangleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2H2v20h20V2ZM6.086 12L7.5 10.586l3 3l6-6L17.915 9L10.5 16.414L6.086 12Z"/>`),
		g.Group(children),
	)
}