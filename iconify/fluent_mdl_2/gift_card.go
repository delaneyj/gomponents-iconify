package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1936 256q23 0 43 9t36 24t24 35t9 44v1184q0 23-9 43t-24 36t-35 24t-44 9H112q-23 0-43-9t-36-24t-24-35t-9-44V368q0-23 9-43t24-36t35-24t44-9h1824zm-16 128H640v139q33-11 64-11q40 0 75 15t61 41t41 61t15 75q0 31-11 64h1035V384zM384 704q0 26 19 45t45 19h64v-64q0-26-19-45t-45-19q-26 0-45 19t-19 45zm320 64q26 0 45-19t19-45q0-26-19-45t-45-19q-26 0-45 19t-19 45v64h64zM128 384v384h139q-11-33-11-64q0-40 15-75t41-61t61-41t75-15q31 0 64 11V384H128zm0 1152h384V987l-147 146l-90-90l146-147H128v640zm1792 0V896H731l146 147l-90 90l-147-146v549h1280z"/>`),
		g.Group(children),
	)
}