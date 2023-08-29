package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareback(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-448-756q-19-20-31-6l-218 232q-7 7-7 18.5t7 19.5l218 231q12 13 31-7V268zm320 0q-18-20-31-6l-218 232q-7 7-7 18.5t7 19.5l218 231q13 13 31-7V268z"/>`),
		g.Group(children),
	)
}