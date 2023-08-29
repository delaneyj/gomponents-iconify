package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckListCheckMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1728 1621l211-210l90 90l-301 301l-173-173l90-90l83 82zm301-1272l-301 301l-173-173l90-90l83 82l211-210l90 90zm-301 504l211-210l90 90l-301 301l-173-173l90-90l83 82zm0 384l211-210l90 90l-301 301l-173-173l90-90l83 82z"/>`),
		g.Group(children),
	)
}