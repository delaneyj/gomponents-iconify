package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteColumns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 0h640v1092l-320 319l-320-320V0zm1280 0v1024h-512V896h384V0h128zM128 896h384v128H0V0h128v896zm1149 606l-226 226l226 227l-90 90l-227-226l-227 227l-90-91l227-227l-227-227l90-90l227 227l227-227l90 91z"/>`),
		g.Group(children),
	)
}