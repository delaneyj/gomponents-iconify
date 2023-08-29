package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowserTabScreenshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1920h384v128H128V0h896v256h768v640h-128V384H896V128H256v1792zm1792-768v896H768v-896h293l128-128h438l128 128h293zm-128 128h-219l-128-128h-330l-128 128H896v640h1024v-640zm-512 0q53 0 99 20t82 55t55 81t20 100q0 53-20 99t-55 82t-81 55t-100 20q-53 0-99-20t-82-55t-55-81t-20-100q0-53 20-99t55-82t81-55t100-20zm0 384q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10z"/>`),
		g.Group(children),
	)
}