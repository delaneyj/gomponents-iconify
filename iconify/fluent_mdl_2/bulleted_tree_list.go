package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletedTreeList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M512 384h1536v128H512V384zm512 640V896h1024v128H1024zm0 512v-128h1024v128H1024zM0 640V256h384v384H0zm128-256v128h128V384H128zm384 768V768h384v384H512zm128-256v128h128V896H640zm-128 768v-384h384v384H512zm128-256v128h128v-128H640z"/>`),
		g.Group(children),
	)
}