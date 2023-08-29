package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinancialMirroredSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2048 1898l-384-384v534h384v-150zm-512-512l-384-384v1046h384v-662zM256 896H128v1152h384v-918L256 896zm448 426l-64-64v790h384V1002l-320 320zM0 256v512h128V475l576 575l384-384l960 961v-182l-960-959l-384 384l-485-486h293V256H0z"/>`),
		g.Group(children),
	)
}