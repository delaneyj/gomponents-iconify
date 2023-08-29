package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m388.632 393.82l107.191-137.88l-107.139-137.762l-25.26 19.644l91.864 118.122l-91.92 118.236l25.264 19.64zm-240.053-19.639L56.712 255.999l91.917-118.176l-25.258-19.646L16.177 255.993l107.137 137.826l25.265-19.638zM330.529 16h-32.97L178.441 496h32.971L330.529 16z"/>`),
		g.Group(children),
	)
}