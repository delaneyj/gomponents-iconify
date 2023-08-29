package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 17.175L7.4 4.6L10 2h8q.825 0 1.413.588T20 4v13.175Zm.5 6.125L15.2 18l1.425-1.4L20 19.975V20q0 .825-.588 1.413T18 22H6q-.825 0-1.413-.588T4 20V8l.6-.6L.7 3.5l1.425-1.4L21.9 21.875L20.5 23.3Z"/>`),
		g.Group(children),
	)
}