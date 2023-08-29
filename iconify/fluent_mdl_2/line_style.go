package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 384v128H0V384h1920zM128 896H0V768h128v128zm256 0H256V768h128v128zm128-128h128v128H512V768zm384 128H768V768h128v128zm256 0h-128V768h128v128zm128-128h128v128h-128V768zm384 128h-128V768h128v128zm128-128h128v128h-128V768zM256 1280H0v-128h256v128zm128-128h256v128H384v-128zm384 0h256v128H768v-128zm384 0h256v128h-256v-128zm384 0h256v128h-256v-128zM0 1536h384v128H0v-128zm512 0h128v128H512v-128zm256 0h384v128H768v-128zm768 0h384v128h-384v-128zm-256 0h128v128h-128v-128z"/>`),
		g.Group(children),
	)
}