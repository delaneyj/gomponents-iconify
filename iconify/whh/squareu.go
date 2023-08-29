package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-736q0-13-9-22.5t-22.5-9.5t-23 9.5t-9.5 22.5v352q0 27-19 45.5t-45 18.5h-128q-26 0-45-19t-19-45V288q0-13-9-22.5t-22.5-9.5t-23 9.5t-9.5 22.5v352q0 53 37.5 90.5t90.5 37.5h128q53 0 90.5-37.5t37.5-90.5V288z"/>`),
		g.Group(children),
	)
}