package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraVideoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2.72l1.515-.38A2 2 0 0 1 22 9.28v5.44a2 2 0 0 1-2.485 1.94L18 16.28V17a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V5Zm16 9.22l2 .5V9.28l-2 .5v4.44Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}