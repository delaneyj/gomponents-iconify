package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squares(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-320-768h-128q-53 0-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5h128q26 0 45 19t19 45v64q0 27-19 45.5t-45 18.5h-128q-27 0-45.5-18.5t-18.5-45.5q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5q0 53 37.5 90.5t90.5 37.5h128q53 0 90.5-37.5t37.5-90.5v-64q0-53-37.5-90.5t-90.5-37.5h-128q-26 0-45-18.5t-19-45t19-45.5t45-19h128q27 0 45.5 19t18.5 45q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5q0-53-37.5-90.5t-90.5-37.5z"/>`),
		g.Group(children),
	)
}