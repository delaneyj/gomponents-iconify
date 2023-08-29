package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareseven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-320-768h-224q-13 0-22.5 9.5t-9.5 22.5t9.5 22.5t22.5 9.5h224q27 0 45.5 19t18.5 45q4 43-16 80q-33 61-128 240q-16 22-16 32q0 12 10 22t22 10q13 0 22.5-9.5t9.5-22.5l144-272q16-30 16-80q0-53-37.5-90.5t-90.5-37.5z"/>`),
		g.Group(children),
	)
}