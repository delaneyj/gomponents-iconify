package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Style(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm64-848q0-46-33-79t-79-33h-672q-46 0-79 33t-33 79v158l.5.5l1.5 1.5q-2 9-2 16q0 40 28 68t68 28q96 0 96-98q-1-40-29-67q-3-16-3-27q0-21 30.5-42.5t65.5-21.5q36 0 66 20.5t30 43.5q0 13-2 27q-30 28-30 69q0 40 28 68t68 28q30 0 54-16.5t35-43.5q38 10 69 49t49 89q-21 13-34 34.5t-13 47.5q0 40 28 68t68 28t68-28t28-68q0-26-12.5-47.5t-34.5-34.5q25-60 74-101t101-41V176z"/>`),
		g.Group(children),
	)
}