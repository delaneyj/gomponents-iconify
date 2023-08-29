package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplayLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M132.56 156.1a6 6 0 0 0-9.11 0l-48 56A6 6 0 0 0 80 222h96a6 6 0 0 0 4.56-9.9ZM93.05 210l35-40.78L163 210ZM230 64v112a22 22 0 0 1-22 22h-16a6 6 0 0 1 0-12h16a10 10 0 0 0 10-10V64a10 10 0 0 0-10-10H48a10 10 0 0 0-10 10v112a10 10 0 0 0 10 10h16a6 6 0 0 1 0 12H48a22 22 0 0 1-22-22V64a22 22 0 0 1 22-22h160a22 22 0 0 1 22 22Z"/>`),
		g.Group(children),
	)
}