package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.93 3.362c.31-.81-.484-1.604-1.293-1.293L2.64 5.915c-.906.348-.834 1.653.105 1.9l4.024 1.06a.5.5 0 0 1 .357.356l1.059 4.024c.247.94 1.552 1.01 1.9.105l3.846-9.998Zm-.934-.36l-3.845 9.999l-1.06-4.025a1.5 1.5 0 0 0-1.068-1.069L2.998 6.848l9.998-3.845Z"/>`),
		g.Group(children),
	)
}