package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 256h-64l-64 64l-192-64H192l384 128l64 64l-64 64l-384-128l-64-64l64-64H64q-26 0-45-19T0 192V64q0-27 19-45.5T64 0h640q27 0 45.5 18.5T768 64v128q0 26-18.5 45T704 256zM576 576l64 64l-64 64l-384-128l-64-64l64-64zm0 192l64 64l-64 64l-384-128l-64-64l64-64zm0 192l-64 64H384l-192-64l-64-64l64-64z"/>`),
		g.Group(children),
	)
}