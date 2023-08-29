package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareprevious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-448-756q-18-20-31-6l-218 231q-7 8-7 19.5t7 18.5l218 232q12 13 31-7V268zm320 52q0-26-19-45t-45-19h-64q-27 0-45.5 19t-18.5 45v384q0 27 18.5 45.5t45.5 18.5h64q26 0 45-18.5t19-45.5V320z"/>`),
		g.Group(children),
	)
}