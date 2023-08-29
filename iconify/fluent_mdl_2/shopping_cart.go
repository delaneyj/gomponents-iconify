package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1600 1536q40 0 75 15t61 41t41 61t15 75q0 40-15 75t-41 61t-61 41t-75 15q-40 0-75-15t-61-41t-41-61t-15-75q0-31 11-64H885q11 33 11 64q0 40-15 75t-41 61t-61 41t-75 15q-40 0-75-15t-61-41t-41-61t-15-75q0-55 29-102t80-71L189 256H0V128h281l85 256h1682l-298 896H665l85 256h850zM409 512l213 640h1035l213-640H409zm359 1216q0-26-19-45t-45-19q-26 0-45 19t-19 45q0 26 19 45t45 19q26 0 45-19t19-45zm832 64q26 0 45-19t19-45q0-26-19-45t-45-19q-26 0-45 19t-19 45q0 26 19 45t45 19z"/>`),
		g.Group(children),
	)
}