package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1024h-384v512h-512v-512H896v768H384v-768H0V896h384V128h512v768h256V384h512v512h384v128zM768 256H512v1408h256V256zm768 256h-256v896h256V512z"/>`),
		g.Group(children),
	)
}