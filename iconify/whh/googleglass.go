package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Googleglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m704 1024l-63-84q-87 20-193 20q-122 0-225-34.5t-163-93T0 704q0-114 31.5-258T117 170l50 18q-48 118-75.5 248T64 668q0 46 20 84t55.5 64t80.5 43t100 27v-22q0-13 9.5-22.5T352 832t22.5 9.5T384 864v29q34 3 64 3q32 0 64-2v-30q0-13 9.5-22.5T544 832t22.5 9.5T576 864v22q65-10 118-29q73-38 105.5-81T832 668q0-87-20.5-197T754 254q-21-5-44.5-29.5T668 163q-7-14-12-25t-10.5-28.5T639 79t3-28t14.5-25T687 8.5T736 0q78 160 119 390.5T896 868q0 21-12.5 47t-35 51t-61 41.5T704 1024z"/>`),
		g.Group(children),
	)
}