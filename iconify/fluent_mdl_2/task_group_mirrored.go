package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskGroupMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256v896h2048V256H0zm128 128h1792v640H128V384zm128 256h1024V512H256v128zm384 256h640V768H640v128zm768 0h384V512h-384v384zm128-256h128v128h-128V640zm256 640H256v128h1536v-128zm-256 256H512v128h1024v-128z"/>`),
		g.Group(children),
	)
}