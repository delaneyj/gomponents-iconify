package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24 10l-1.414 1.414L26.172 15H5.828l3.586-3.586L8 10l-6 6l6 6l1.414-1.414L5.828 17h20.344l-3.586 3.586L24 22l6-6l-6-6z"/>`),
		g.Group(children),
	)
}