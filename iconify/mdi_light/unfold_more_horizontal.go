package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldMoreHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.74 8.83l-.7.71L11.5 6L7.96 9.54l-.7-.71l4.24-4.24l4.24 4.24m0 7.34l-4.24 4.24l-4.24-4.24l.7-.71L11.5 19l3.54-3.54l.7.71Z"/>`),
		g.Group(children),
	)
}