package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 22h14v2H16zm-2-2.6L12.6 18L6 24.6L3.4 22L2 23.4l4 4zM16 8h14v2H16zm-2-2.6L12.6 4L6 10.6L3.4 8L2 9.4l4 4z"/>`),
		g.Group(children),
	)
}