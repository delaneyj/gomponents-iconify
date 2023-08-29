package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpPhotoCameraBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.83 5L15 3H9L7.17 5H2v16h20V5h-5.17zM6 17l3-4l2.25 3l3-4L18 17H6z"/>`),
		g.Group(children),
	)
}