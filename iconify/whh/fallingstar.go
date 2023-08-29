package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fallingstar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m401 602l-14-5l-31-67q50-124 159.5-237.5t243.5-190T1024 0v192q-176 35-347 146T401 602zm107 140L384 861v145q-2 9-9.5 14.5T358 1023l-121-75l-149 50q-10 1-17.5-5.5T62 975l30-153L2 697q-4-10 0-19t13-13l152-20l90-127q7-7 17-6t17 9l64 140l146 47q8 6 10 16t-3 18z"/>`),
		g.Group(children),
	)
}