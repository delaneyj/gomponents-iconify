package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CustomList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 0h1536v2048H256V0zm1408 1920V128H384v1792h1280zM1536 384v128H768V384h768zm-128 384v128H768V768h640zm-256 384v128H768v-128h384zm384 384v128H768v-128h768zM640 384v128H512V384h128zm0 384v128H512V768h128zm0 384v128H512v-128h128zm0 384v128H512v-128h128z"/>`),
		g.Group(children),
	)
}