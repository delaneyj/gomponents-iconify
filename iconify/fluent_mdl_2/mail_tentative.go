package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailTentative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v640h-128V648l-896 447l-896-447v888h1408v128H0V384h2048zM1024 953l881-441H143l881 441zm640 967h128v128h-128v-128zm64-768q53 0 99 20t82 55t55 81t20 100q0 46-14 81t-35 63t-47 50t-46 45t-36 45t-14 52v48h-128v-48q0-47 14-81t35-63t47-50t46-45t36-45t14-52q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50h-128q0-53 20-99t55-82t81-55t100-20z"/>`),
		g.Group(children),
	)
}