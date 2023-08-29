package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldLessVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.71 16.74l-4.25-4.24l4.25-4.24l.7.7l-3.53 3.54l3.53 3.54l-.7.7m-14.42 0l-.7-.7l3.53-3.54l-3.53-3.54l.7-.7l4.25 4.24l-4.25 4.24Z"/>`),
		g.Group(children),
	)
}