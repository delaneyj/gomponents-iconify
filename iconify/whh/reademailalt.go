package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reademailalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.428 1024h-896q-27 0-45.5-18.5T.428 960V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v832q0 27-18.5 45.5t-45.5 18.5zm-64-768q0-53-37.5-90.5t-90.5-37.5h-512q-53 0-90.5 37.5t-37.5 90.5v384q0 26 19 45t45 19h32q21 0 39 8.5t25 23.5q5 17 16.5 66t15.5 62q8 20 24 26t45 6h249q29 0 43.5-6t22.5-26q19-90 32-128q13-32 64-32h32q27 0 45.5-19t18.5-45V256zm-385 372q-1 12-9 19l-45 45q-10 10-25 9q-15 1-25-9l-45-45q-5-5-7-11l-90-88q-9-9-9-22.5t9-22.5l46-46q9-9 22.5-9t23.5 9l71 71l239-263q10-9 23.5-9t22.5 9l46 46q9 9 9 22.5t-9 22.5z"/>`),
		g.Group(children),
	)
}