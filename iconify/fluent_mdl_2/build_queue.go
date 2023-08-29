package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildQueue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1024h-512v512h-512v512H0V1024h512V512h512V0h1024v1024zM896 1152H128v768h768v-768zm512-512H640v384h384v384h384V640zm512-512h-768v384h384v384h384V128z"/>`),
		g.Group(children),
	)
}