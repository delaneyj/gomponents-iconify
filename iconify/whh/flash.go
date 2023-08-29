package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-480-768h-192q-13 0-22.5 9.5t-9.5 22.5v448q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V592q0-7 4.5-11.5t11.5-4.5h80q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5h-80q-7 0-11.5-4.5t-4.5-11.5V336q0-7 4.5-11.5t11.5-4.5h144q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5zm384 448h-160V288q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5v448q0 13 9.5 22.5t22.5 9.5h192q13 0 22.5-9.5t9.5-22.5t-9.5-22.5t-22.5-9.5z"/>`),
		g.Group(children),
	)
}