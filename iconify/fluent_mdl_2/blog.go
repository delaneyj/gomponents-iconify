package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v1536H731l-475 475v-475H0V0h2048zM128 128v256h1792V128H128zm1792 1280V512H128v896h256v293l293-293h1243zm-640-768h512v640h-512V640zm128 512h256V768h-256v384zM256 768h896v128H256V768zm0 256h896v128H256v-128z"/>`),
		g.Group(children),
	)
}