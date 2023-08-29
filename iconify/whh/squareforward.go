package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-391-531l-218-231q-13-14-31 6v488q18 20 31 7l218-232q7-7 7-18.5t-7-19.5zm320 0l-218-231q-12-14-31 6v488q19 20 31 7l218-232q7-7 7-18.5t-7-19.5z"/>`),
		g.Group(children),
	)
}