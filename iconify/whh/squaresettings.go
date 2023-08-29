package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squaresettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-128-512q0-23-6-51l68-65q7-6-3-23l-44-67q-10-16-21-14l-85 26q-36-31-77-46l-17-70q-3-10-24-10h-85q-21 0-24 10l-19 67q-45 15-84 49l-85-26q-11-2-21 14l-44 67q-10 17-2 23l67 65q-6 28-6 51t6 51l-67 65q-8 6 2 23l44 67q10 16 21 14l85-26q36 31 77 46l17 70q3 10 24 10h85q21 0 24-10l19-67q45-15 84-49l85 26q11 2 21-14l44-67q10-17 3-23l-68-65q6-28 6-51zm-256 128q-53 0-90.5-37.5t-37.5-90.5t37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5t-37.5 90.5t-90.5 37.5z"/>`),
		g.Group(children),
	)
}