package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarenine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-640q0-53-37.5-90.5t-90.5-37.5h-128q-53 0-90.5 37.5t-37.5 90.5v64q0 53 37.5 90.5t90.5 37.5h128q34 0 64-17v81q0 26-19 45t-45 19h-128q-26 0-45-19t-19-45q0-13-9.5-22.5t-22.5-9.5t-22.5 9.5t-9.5 22.5q0 53 37.5 90.5t90.5 37.5h128q53 0 90.5-37.5t37.5-90.5V384zm-128 128h-128q-26 0-45-19t-19-45v-64q0-27 19-45.5t45-18.5h128q27 0 45.5 18.5t18.5 45.5v64q0 26-18.5 45t-45.5 19z"/>`),
		g.Group(children),
	)
}