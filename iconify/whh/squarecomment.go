package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Squarecomment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-384-832q-87 0-160.5 38.5t-116.5 105t-43 144.5q0 59 25.5 112.5t70.5 93.5q-3 37-12.5 80t-19.5 66h1q34 0 80.5-21.5t76.5-56.5q48 14 98 14q87 0 160.5-38.5t116.5-105t43-144.5t-43-144.5t-116.5-105t-160.5-38.5z"/>`),
		g.Group(children),
	)
}