package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCheckedMirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 19.4L28.6 18L22 24.6L19.4 22L18 23.4l4 4zM2 22h14v2H2zM30 5.4L28.6 4L22 10.6L19.4 8L18 9.4l4 4zM2 8h14v2H2z"/>`),
		g.Group(children),
	)
}