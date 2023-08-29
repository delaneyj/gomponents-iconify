package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarefive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-448q0-53-37.5-90.5t-90.5-37.5h-176q-7 0-11.5-4.5t-4.5-11.5v-96q0-6 4.5-11t11.5-5h272q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-320q-13 0-22.5 9.5t-9.5 22.5v160q0 27 18.5 45.5t45.5 18.5h192q27 0 45.5 19t18.5 45v64q0 27-18.5 45.5t-45.5 18.5h-128q-27 0-45.5-18.5t-18.5-45.5q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5q0 53 37.5 90.5t90.5 37.5h128q53 0 90.5-37.5t37.5-90.5v-64z"/>`),
		g.Group(children),
	)
}