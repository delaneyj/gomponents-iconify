package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayPassOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-.825 0-1.413.588T10 22H5V2h5q0 .825.588 1.413T12 4q.825 0 1.413-.588T14 2h5v20h-5q0-.825-.588-1.413T12 20Zm0-2q1.075 0 1.988.537T15.45 20H17V4h-1.55q-.55.925-1.463 1.463T12 6q-1.075 0-1.988-.537T8.55 4H7v16h1.55q.55-.925 1.463-1.463T12 18Zm0-6Z"/>`),
		g.Group(children),
	)
}