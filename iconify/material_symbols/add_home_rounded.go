package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddHomeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 16.5v2q0 .2.15.35T18 19q.2 0 .35-.15t.15-.35v-2h2q.2 0 .35-.15T21 16q0-.2-.15-.35t-.35-.15h-2v-2q0-.2-.15-.35T18 13q-.2 0-.35.15t-.15.35v2h-2q-.2 0-.35.15T15 16q0 .2.15.35t.35.15h2ZM18 21q-2.075 0-3.538-1.463T13 16q0-2.075 1.463-3.538T18 11q2.075 0 3.538 1.463T23 16q0 2.075-1.463 3.538T18 21ZM4 17V8q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 8v1.3q-.5-.15-1-.225T18 9q-2.925 0-4.963 2.038T11 16q0 .8.163 1.55t.512 1.45H6q-.825 0-1.413-.588T4 17Z"/>`),
		g.Group(children),
	)
}