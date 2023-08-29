package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Closet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 10.551v-.678A4.005 4.005 0 0 0 16 6c0-2.206-1.794-4-4-4S8 3.794 8 6h2c0-1.103.897-2 2-2s2 .897 2 2s-.897 2-2 2a1 1 0 0 0-1 1v1.551l-8.665 7.702A1.001 1.001 0 0 0 3 20h18a1.001 1.001 0 0 0 .664-1.748L13 10.551zM5.63 18L12 12.338L18.37 18H5.63z"/>`),
		g.Group(children),
	)
}