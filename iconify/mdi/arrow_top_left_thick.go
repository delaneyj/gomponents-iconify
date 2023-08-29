package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeftThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.12 11.94v4.95H5.69V5.69h11.2v3.43h-4.95l6.37 6.38l-2.81 2.81l-6.38-6.37Z"/>`),
		g.Group(children),
	)
}