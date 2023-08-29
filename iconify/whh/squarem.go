package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-735q0-14-10-23.5t-22-9.5q-16 0-27 16l-133 186l-132-186q-12-16-28-16q-13 0-22.5 9.5t-9.5 23.5v447q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V388l100 140q11 15 28 15q16 0 28-15l100-140v348q0 13 9.5 22.5t22.5 9.5t22.5-9.5t9.5-22.5V289z"/>`),
		g.Group(children),
	)
}