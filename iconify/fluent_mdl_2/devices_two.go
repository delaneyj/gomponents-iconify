package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1663q0 27-10 50t-27 41t-41 28t-50 10H896q-27 0-50-10t-40-27t-28-41t-10-50v-128H122q-25 0-47-9t-39-26t-26-39t-10-48q0-35 11-70t36-61l209-220V256h1408v512h256q26 0 49 10t40 27t28 40t11 49v769zM768 1408v-256H348l-206 217q-6 7-9 18t-5 21h640zm0-512q0-27 10-50t27-40t41-28t50-10h640V384H384v640h384V896zm1152 768V896H896v768h1024zm-640-256h256v128h-256v-128z"/>`),
		g.Group(children),
	)
}