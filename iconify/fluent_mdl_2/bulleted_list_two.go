package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletedListTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v128H512V384h1536zM512 896h1536v128H512V896zm0 512h1536v128H512v-128zM0 256h384v384H0V256zm128 256h128V384H128v128zM0 768h384v384H0V768zm128 256h128V896H128v128zM0 1280h384v384H0v-384zm128 256h128v-128H128v128z"/>`),
		g.Group(children),
	)
}