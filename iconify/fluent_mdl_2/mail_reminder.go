package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReminder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v640h-128V648l-896 447l-896-447v888h896v128H0V384h2048zM1024 953l881-441H143l881 441zm576 199q66 0 124 25t101 69t69 102t26 124v192h128v128h-256v64q0 40-15 75t-41 61t-61 41t-75 15q-40 0-75-15t-61-41t-41-61t-15-75v-64h-256v-128h128v-192q0-66 25-124t68-101t102-69t125-26zm64 640h-128v64q0 26 19 45t45 19q26 0 45-19t19-45v-64zm128-128v-192q0-40-15-75t-41-61t-61-41t-75-15q-40 0-75 15t-61 41t-41 61t-15 75v192h384z"/>`),
		g.Group(children),
	)
}