package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarequora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-128-544q0-119-75-203.5t-181-84.5t-181 84.5t-75 203.5t75 203.5t181 84.5q63 0 118-33q21 14 52 23.5t54 9.5q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5q-24 0-44-19q76-85 76-205zm-224 96q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5q24 0 44 19q-34 45-76 45q-53 0-90.5-65.5t-37.5-158.5t37.5-158.5t90.5-65.5t90.5 65.5t37.5 158.5q0 60-18 114q-42-18-78-18z"/>`),
		g.Group(children),
	)
}