package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Money(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M768 896h384v-96h-128V352H910L762 489l77 80q42-37 55-57h2v288H768v96zm512-256q0 70-21 142t-59.5 134t-101.5 101t-138 39t-138-39t-101.5-101T661 782t-21-142t21-142t59.5-134T822 263t138-39t138 39t101.5 101t59.5 134t21 142zm512 256V384q-106 0-181-75t-75-181H384q0 106-75 181t-181 75v512q106 0 181 75t75 181h1152q0-106 75-181t181-75zm128-832v1152q0 26-19 45t-45 19H64q-26 0-45-19t-19-45V64q0-26 19-45T64 0h1792q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}