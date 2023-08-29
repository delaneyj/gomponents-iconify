package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBrandedColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1664H0V128zm768 640v384h512V768H768zm512-128V256H768v384h512zm-512 640v384h512v-384H768z"/>`),
		g.Group(children),
	)
}