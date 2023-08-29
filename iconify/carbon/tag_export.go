package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m31 18l-6-6l-1.414 1.414L27.172 17H17v2h10.172l-3.586 3.586L25 24l6-6zm-21-4a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 1.998 2.004A2.002 2.002 0 0 0 10 8z"/><path fill="currentColor" d="m21.06 25l-3.002 3L4 13.941V4h9.942l7 7l1.414-1.414l-7.001-7A2 2 0 0 0 13.941 2H4a2 2 0 0 0-2 2v9.941a2 2 0 0 0 .586 1.414l14.058 14.06a2.001 2.001 0 0 0 2.828 0l3.002-3Z"/>`),
		g.Group(children),
	)
}