package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dailybooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M232 462q83 0 145-51l81 29l-34-81q38-59 38-127q0-95-67.5-162.5T232 2T69.5 69.5T2 232t67.5 162.5T232 462zM105 172l47-11q24-44 73-44q29 0 48 15l47-11q8-1 15.5 4t9.5 14l31 134q5 22-14 27l-216 50q-19 5-25-17L90 199q-5-21 15-27zm132 120q24 0 41-17t17-41t-17-41t-41-17t-41 17t-17 41t17 41t41 17zm0-95q15 0 26 10.5t11 26.5q0 15-11 26t-26 11q-16 0-26.5-11T200 234q0-16 10.5-26.5T237 197z"/>`),
		g.Group(children),
	)
}