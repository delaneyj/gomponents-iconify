package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 0v1792H256V0h1536zm-128 128H384v1536h1280V128zM640 896H512V768h128v128zm896 0H768V768h768v128zm-896 384H512v-128h128v128zm896 0H768v-128h768v128zM640 512H512V384h128v128zm896 0H768V384h768v128z"/>`),
		g.Group(children),
	)
}