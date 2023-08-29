package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightMappedShapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 640v640h-896v-256H979l-339 226v158h384v640H128v-640h384v-158L77 960l435-290V512H384q-53 0-99-20t-82-55t-55-81t-20-100q0-53 20-99t55-82t81-55T384 0h384q53 0 99 20t82 55t55 81t20 100q0 53-20 99t-55 82t-81 55t-100 20H640v158l339 226h173V640h896zM768 384q27 0 50-10t40-27t28-41t10-50q0-27-10-50t-27-40t-41-28t-50-10H384q-27 0-50 10t-40 27t-28 41t-10 50q0 27 10 50t27 40t41 28t50 10h384zM384 1664v128h384v-128H384zm461-704L576 781L307 960l269 179l269-179zm947-64h-384v128h384V896z"/>`),
		g.Group(children),
	)
}