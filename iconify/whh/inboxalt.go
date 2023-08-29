package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inboxalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.428 1024h-896q-26 0-45-19t-19-45V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v832q0 26-18.5 45t-45.5 19zm-64-768q0-53-37.5-90.5t-90.5-37.5h-512q-53 0-90.5 37.5t-37.5 90.5v384q0 27 18.5 45.5t45.5 18.5h32q20 0 38.5 8.5t25.5 23.5q5 17 16.5 66t15.5 62q8 20 23.5 26t45.5 6h249q29 0 43.5-6t22.5-26q18-90 32-128q13-32 64-32h32q26 0 45-18.5t19-45.5V256zm-352 434q-13 14-32 14t-33-14l-214-217q-12-13-8-19t19-6h108V320q0-27 18.5-45.5t45.5-18.5h128q27 0 45.5 18.5t18.5 45.5v128h107q15 0 19 6t-8 19z"/>`),
		g.Group(children),
	)
}