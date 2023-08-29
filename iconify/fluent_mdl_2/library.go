package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Library(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 1792V256h384v1536H640zM768 384v1280h128V384H768zM128 1792V256h384v1536H128zM256 384v1280h128V384H256zm1235-151l484 1450l-346 116l-484-1450l346-116zm-204 186l412 1238l134-44l-413-1238l-133 44z"/>`),
		g.Group(children),
	)
}