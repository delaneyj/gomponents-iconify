package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteRows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 128H0V0h1024v512H896V128zm515 832l-320 320H0V640h1092l319 320zm-515 448h128v512H0v-128h896v-384zm1059-131l-227-227l-227 227l-90-90l227-227l-227-227l91-90l226 226l227-226l90 90l-226 227l227 227l-91 90z"/>`),
		g.Group(children),
	)
}