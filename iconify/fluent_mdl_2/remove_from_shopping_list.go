package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveFromShoppingList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 768v128H512V768h128zm896 0v128H768V768h768zM512 1280v-128h128v128H512zm256 0v-128h768v128H768zM640 384v128H512V384h128zm896 0v128H768V384h768zM384 128v1536h896v128H256V0h1536v1280h-128V128H384zm1645 1389l-211 211l211 211l-90 90l-211-211l-211 211l-90-90l211-211l-211-211l90-90l211 211l211-211l90 90z"/>`),
		g.Group(children),
	)
}