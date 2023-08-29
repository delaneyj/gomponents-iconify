package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Constellation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m951 189l-47 332q25 11 40.5 34.5T960 608q0 40-28 68t-68 28q-29 0-54-16L505 893q7 17 7 35q0 40-28 68t-68 28t-68-28t-28-68q0-18 7-35L131 697q-17 7-35 7q-40 0-68-28T0 608t28-68t68-28q22 0 43 10l186-204q-5-16-5-30q0-40 28-68t68-28q44 0 73 35L833 89q2-37 29.5-63T928 0q40 0 68 28t28 68q0 33-20.5 59T951 189zM365 369L183 568q9 20 9 40q0 24-13 47l190 190q23-13 47-13q25 0 48 13l309-208q-5-15-5-29q0-11 3-23L480 359q-28 25-64 25q-27 0-51-15zm146-66l295 229q15-12 35-17l47-332q-25-12-41-35L511 282q0 1 .5 3t.5 3q0 5-1 15z"/>`),
		g.Group(children),
	)
}