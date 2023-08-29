package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenericScanFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 128h512v512h-128V256h-384V128zM128 256v384H0V128h512v128H128zm1792 1536v-384h128v512h-512v-128h384zM256 1664V384h1536v1280H256zM1536 640v768h128V640h-128zm-384 0v768h128V640h-128zm-384 0v768h128V640H768zm-384 0v768h128V640H384zm-256 768v384h384v128H0v-512h128z"/>`),
		g.Group(children),
	)
}