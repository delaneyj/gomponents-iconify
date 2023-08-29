package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1568 256h-217L882 1664h366l-128 128H416l128-128h161l469-1408H864l128-128h704l-128 128z"/>`),
		g.Group(children),
	)
}