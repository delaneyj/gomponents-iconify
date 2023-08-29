package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegistryEditor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 896h512v1152H0V384h1152v512zM640 512v384h384V512H640zm384 896v-384H640v384h384zM128 512v384h384V512H128zm0 512v384h384v-384H128zm384 896v-384H128v384h384zm512 0v-384H640v384h384zm512 0v-384h-384v384h384zm-384-512h384v-384h-384v384zm832-960l-384 384l-384-384l384-384l384 384zm-384-239l-239 239l239 239l239-239l-239-239z"/>`),
		g.Group(children),
	)
}