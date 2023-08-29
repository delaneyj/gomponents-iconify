package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squareeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-448q0-58-43-96q43-38 43-96q0-53-37.5-90.5t-90.5-37.5h-128q-53 0-90.5 37.5t-37.5 90.5q0 58 43 96q-43 38-43 96v64q0 53 37.5 90.5t90.5 37.5h128q53 0 90.5-37.5t37.5-90.5v-64zm-128 128h-128q-27 0-45.5-18.5t-18.5-45.5v-64q0-26 18.5-45t45.5-19h128q27 0 45.5 19t18.5 45v64q0 27-18.5 45.5t-45.5 18.5zm0-256h-128q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h128q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5z"/>`),
		g.Group(children),
	)
}