package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PulseTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m8.471 7.237l3.056 12.992c.233.993 1.63 1.04 1.929.065l2.945-9.58l.384 1.527a1 1 0 0 0 .97.756H21a1 1 0 1 0 0-2h-2.466l-1.069-4.241c-.247-.981-1.628-1.017-1.925-.05l-2.912 9.472L9.475 2.771c-.24-1.02-1.688-1.03-1.943-.014l-2.065 8.24H3a1 1 0 0 0 0 2h3.247a1 1 0 0 0 .97-.757l1.254-5.003Z"/>`),
		g.Group(children),
	)
}