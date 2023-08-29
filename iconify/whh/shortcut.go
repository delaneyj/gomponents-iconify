package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shortcut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-832q0-26-19-45t-45-19h-448q-16 0-32 8t-24 16l-8 8l143 144q-64 35-122.5 84t-107 108.5t-77 130.5t-28.5 141q0 7-.5 20.5t-.5 24t1 19.5q0 27 18.5 45.5t45 18.5t45.5-18.5t19-45.5q19-85 85-148t159-96.5t206-41.5l158 158q1 0 8.5-8t15.5-24t8-32V192z"/>`),
		g.Group(children),
	)
}