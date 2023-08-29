package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1137 972l306-297l-422-62l-189-382l-189 382l-422 62l306 297l-73 421l378-199l377 199zm527-357q0 22-26 48l-363 354l86 500q1 7 1 20q0 50-41 50q-19 0-40-12l-449-236l-449 236q-22 12-40 12q-21 0-31.5-14.5T301 1537q0-6 2-20l86-500L25 663Q0 636 0 615q0-37 56-46l502-73L783 41q19-41 49-41t49 41l225 455l502 73q56 9 56 46z"/>`),
		g.Group(children),
	)
}