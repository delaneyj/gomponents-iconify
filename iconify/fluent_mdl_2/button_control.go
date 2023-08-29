package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h2048v1152H0V384zm1920 1024V512H128v896h1792zm-128-512v128H768V896h1024zM448 1152q-40 0-75-15t-61-41t-41-61t-15-75q0-40 15-75t41-61t61-41t75-15q40 0 75 15t61 41t41 61t15 75q0 40-15 75t-41 61t-61 41t-75 15zm0-256q-26 0-45 19t-19 45q0 26 19 45t45 19q26 0 45-19t19-45q0-26-19-45t-45-19z"/>`),
		g.Group(children),
	)
}