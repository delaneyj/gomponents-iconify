package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-736q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5v208q0 7-5 11.5t-11 4.5h-224q-7 0-12-4.5t-5-11.5V288q0-13-9-22.5t-22.5-9.5t-23 9.5t-9.5 22.5v448q0 13 9.5 22.5t23 9.5t22.5-9.5t9-22.5V592q0-7 5-11.5t12-4.5h224q6 0 11 4.5t5 11.5v144q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V288z"/>`),
		g.Group(children),
	)
}