package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gradient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-192-256h-64v-64h64v-64h-64v-64h64v-64h-64v-64h64v-64h-64v-64h64v-64h-64v-64h64v-64h-64v64h-64v-64h-64v64h-64v-64h-288q-13 0-22.5 9.5t-9.5 22.5v704q0 13 9.5 22.5t22.5 9.5h224v-64h64v64h64v-64h64v64h64v-64h64v-64zm-128 64v-64h64v64h-64zm0-192h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm-64 0v-64h64v64h-64zm-128 0v-64h64v64h-64zm128 512v64h-64v-64h64zm0-64h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm0-128h64v64h-64v-64zm-64 0v-64h64v64h-64zm64 128h-64v-64h64v64zm0 128h-64v-64h64v64zm0 128h-64v-64h64v64zm-128 64v-64h64v64h-64zm0-128v-64h64v64h-64zm0-128v-64h64v64h-64zm0-128v-64h64v64h-64z"/>`),
		g.Group(children),
	)
}