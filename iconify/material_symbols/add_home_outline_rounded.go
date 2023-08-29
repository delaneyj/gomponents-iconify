package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddHomeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 16.5v2q0 .2.15.35T18 19q.2 0 .35-.15t.15-.35v-2h2q.2 0 .35-.15T21 16q0-.2-.15-.35t-.35-.15h-2v-2q0-.2-.15-.35T18 13q-.2 0-.35.15t-.15.35v2h-2q-.2 0-.35.15T15 16q0 .2.15.35t.35.15h2ZM18 21q-2.075 0-3.538-1.463T13 16q0-2.075 1.463-3.538T18 11q2.075 0 3.538 1.463T23 16q0 2.075-1.463 3.538T18 21ZM4 17V8q0-.475.213-.9t.587-.7l6-4.5q.275-.2.575-.3T12 1.5q.325 0 .625.1t.575.3l6 4.5q.375.275.588.7T20 8v1.3q-.475-.15-.975-.225T18 9V8l-6-4.5L6 8v9h5.075q.075.525.225 1.025t.375.975H6q-.825 0-1.413-.588T4 17Zm8-6.75Z"/>`),
		g.Group(children),
	)
}