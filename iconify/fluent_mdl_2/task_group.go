package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v896H0V256h2048zm-128 128H128v640h1792V384zm-128 256H768V512h1024v128zm-384 256H768V768h640v128zm-768 0H256V512h384v384zM512 640H384v128h128V640zm-256 640h1536v128H256v-128zm256 256h1024v128H512v-128z"/>`),
		g.Group(children),
	)
}