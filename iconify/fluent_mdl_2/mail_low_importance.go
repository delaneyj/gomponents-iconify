package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailLowImportance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h2048v1195l-128-128V648l-896 447l-896-447v888h1333l-128 128H0V384zm143 128l881 441l881-441H143zm1649 1286l147-147l90 90l-301 301l-301-301l90-90l147 147v-774h128v774z"/>`),
		g.Group(children),
	)
}