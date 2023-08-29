package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyPhraseExtraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M896 1152v128H0v-128h896zm640-256v128H640V896h896zM1371 0l549 549v1499H256v-640h128v512h1408V640h-512V128H384v896H256V0h1115zm37 512h293l-293-293v293zM640 1536v-128h896v128H640z"/>`),
		g.Group(children),
	)
}