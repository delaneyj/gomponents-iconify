package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ascending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1216 1024h512l-320 640h320v128h-512l320-640h-320v-128zm-704 614l163-163l90 90l-317 318l-317-318l90-90l163 163V128h128v1510zm811-870l-43 128h-128l256-768h128l256 768h-128l-43-128h-298zm149-448l-107 320h214l-107-320z"/>`),
		g.Group(children),
	)
}