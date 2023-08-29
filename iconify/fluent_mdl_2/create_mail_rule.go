package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreateMailRule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1152v896H640v-896h1408zM935 1280l409 245l409-245H935zm985 640v-591l-576 346l-576-346v591h1152zM928 512q-31 0-54 9t-44 24t-41 31t-45 31t-58 23t-78 10H128v896h384v128H0V384q0-27 10-50t27-40t41-28t50-10h480q45 0 77 9t58 24t46 31t40 31t44 23t55 10h736q27 0 50 10t40 27t28 41t10 50v512h-128V512H928zM128 384v128h480q24 0 42-4t33-13t29-20t32-27q-17-15-31-26t-30-20t-33-13t-42-5H128z"/>`),
		g.Group(children),
	)
}