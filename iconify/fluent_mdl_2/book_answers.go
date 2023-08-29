package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookAnswers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 0h1408v2048H384q-53 0-99-20t-82-55t-55-81t-20-100V256q0-53 20-99t55-82t81-55T384 0zm1280 1920v-256H384q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10h1280zm0-384V128H384q-27 0-50 10t-40 27t-28 41t-10 50v1314q60-34 128-34h1280zm-768-256V640h128v640H896zm0-768V384h128v128H896z"/>`),
		g.Group(children),
	)
}