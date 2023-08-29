package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarej(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-224-768q-13 0-22.5 9.5t-9.5 22.5v352q0 27-19 45.5t-45 18.5h-128q-27 0-45.5-18.5t-18.5-45.5q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5q0 53 37.5 90.5t90.5 37.5h128q53 0 90.5-37.5t37.5-90.5V288q0-13-9.5-22.5t-22.5-9.5z"/>`),
		g.Group(children),
	)
}