package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestBeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1859 1758q14 23 21 47t7 51q0 40-15 75t-41 61t-61 41t-75 15H354q-40 0-75-15t-61-41t-41-61t-15-75q0-27 6-51t21-47l569-992q10-14 10-34V128H640V0h768v128h-128v604q0 19 10 35l569 991zM896 732q0 53-27 99l-331 577h972l-331-577q-27-46-27-99V128H896v604zm799 1188q26 0 44-19t19-45q0-10-2-17t-8-16l-164-287H464l-165 287q-9 15-9 33q0 26 18 45t46 19h1341z"/>`),
		g.Group(children),
	)
}