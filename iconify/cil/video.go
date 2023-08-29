package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m461 132l-101 39.277V72H16v368h344V332.723L461 372h35V132Zm3 206.833l-136-52.89V408H48V104h280v114.057l136-52.89Z"/>`),
		g.Group(children),
	)
}