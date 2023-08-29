package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBrandedRow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1664H0V128zm768 1024h512V768H768v384zM640 768H128v384h512V768zm768 384h512V768h-512v384z"/>`),
		g.Group(children),
	)
}