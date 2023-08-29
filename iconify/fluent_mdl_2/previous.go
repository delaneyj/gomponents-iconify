package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1792V256h128v1536H256zm448-768l1088-768v1536L704 1024zm960 521V503l-738 521l738 521z"/>`),
		g.Group(children),
	)
}