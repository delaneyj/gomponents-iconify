package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IronCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M486 371.316V142.75a399.056 399.056 0 0 1-185.685 69.28a399.135 399.135 0 0 1 70.973-186.002H142.72A399.056 399.056 0 0 1 212 211.714A399.135 399.135 0 0 1 26 140.74v228.57a399.056 399.056 0 0 1 185.685-69.28a399.135 399.135 0 0 1-70.973 185.944H369.28A399.056 399.056 0 0 1 300 300.342a399.135 399.135 0 0 1 186 70.974z"/>`),
		g.Group(children),
	)
}