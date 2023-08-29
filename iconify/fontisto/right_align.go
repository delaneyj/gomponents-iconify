package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightAlign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M27.999 21.333H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm0-5.333H8.444a1.334 1.334 0 0 0-.002 2.666h19.557A1.334 1.334 0 0 0 28.001 16zm0-5.333H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zM1.334 2.666h26.665l.037.001a1.334 1.334 0 1 0 0-2.668L27.997 0h.002H1.334a1.334 1.334 0 0 0-.002 2.666h.002zm26.665 2.667H8.444l-.037-.001a1.334 1.334 0 1 0 0 2.668l.039-.001h-.002h19.555a1.334 1.334 0 0 0 .002-2.666z"/>`),
		g.Group(children),
	)
}