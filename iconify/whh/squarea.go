package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-640q0-53-37.5-90.5t-90.5-37.5h-128q-53 0-90.5 37.5t-37.5 90.5v352q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V592q0-6 4.5-11t11.5-5h224q7 0 11.5 4.5t4.5 11.5v144q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V384zm-80 128h-224q-7 0-11.5-4.5t-4.5-11.5V384q0-26 19-45t45-19h128q27 0 45.5 19t18.5 45v112q0 7-4.5 11.5t-11.5 4.5z"/>`),
		g.Group(children),
	)
}