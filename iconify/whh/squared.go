package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-193-640q0-53-37.5-90.5t-90.5-37.5h-223q-14 0-23 9.5t-9 22.5v448q0 13 9 22.5t23 9.5h223q53 0 90.5-37.5t37.5-90.5V384zm-128 320h-175q-7 0-11.5-4.5t-4.5-11.5V336q0-6 4.5-11t11.5-5h175q27 0 45.5 19t18.5 45v256q0 27-18.5 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}