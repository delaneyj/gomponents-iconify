package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckListCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M493 349L192 650L19 477l90-90l83 82l211-210l90 90zM192 1621l211-210l90 90l-301 301l-173-173l90-90l83 82zm0-768l211-210l90 90l-301 301L19 861l90-90l83 82zm0 384l211-210l90 90l-301 301l-173-173l90-90l83 82z"/>`),
		g.Group(children),
	)
}