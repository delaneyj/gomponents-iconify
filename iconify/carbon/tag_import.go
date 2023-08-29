package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagImport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 18l6-6l1.414 1.414L19.828 17H30v2H19.828l3.586 3.586L22 24l-6-6zm-6-4a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 1.998 2.004A2.002 2.002 0 0 0 10 8z"/><path fill="currentColor" d="m20.059 26l-2.001 2L4 13.941V4h9.942l6 6l1.414-1.414l-6.001-6A2 2 0 0 0 13.941 2H4a2 2 0 0 0-2 2v9.941a2 2 0 0 0 .586 1.414l14.058 14.06a2.001 2.001 0 0 0 2.828 0l2-2.001Z"/>`),
		g.Group(children),
	)
}