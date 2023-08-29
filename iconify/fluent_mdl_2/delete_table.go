package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1819 1728l227 227l-91 90l-227-227l-227 227l-90-90l227-227l-227-227l91-90l226 226l227-226l90 90l-226 227zm226 0l3-3v6l-3-3zm3-1600v1152h-642l-126 126v-126H768v384h512v128H0V128h2048zM640 1280H128v384h512v-384zm0-512H128v384h512V768zm0-512H128v384h512V256zm640 512H768v384h512V768zm0-512H768v384h512V256zm640 512h-512v384h512V768zm0-512h-512v384h512V256z"/>`),
		g.Group(children),
	)
}