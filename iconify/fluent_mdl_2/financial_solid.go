package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinancialSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m0 1898l384-384v534H0v-150zm512-512l384-384v1046H512v-662zm1280-490h128v1152h-384v-918l256-234zm-448 426l64-64v790h-384V1002l320 320zm704-1066v512h-128V475l-576 575l-384-384L0 1627v-182l960-959l384 384l485-486h-293V256h512z"/>`),
		g.Group(children),
	)
}