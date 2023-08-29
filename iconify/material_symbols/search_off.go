package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22q-2.075 0-3.538-1.463T2 17q0-2.075 1.463-3.538T7 12q2.075 0 3.538 1.463T12 17q0 2.075-1.463 3.538T7 22Zm13.6-1l-6.4-6.4q-.3-.325-.637-.662T12.9 13.3q.95-.6 1.525-1.6T15 9.5q0-1.875-1.313-3.188T10.5 5Q8.625 5 7.312 6.313T6 9.5q0 .15.013.288t.037.287q-.45.05-.987.2t-.963.35q-.05-.275-.075-.55T4 9.5q0-2.725 1.888-4.612T10.5 3q2.725 0 4.612 1.888T17 9.5q0 1.075-.338 2.038t-.937 1.762L22 19.6L20.6 21ZM5.225 19.475L7 17.7l1.75 1.775l.725-.7L7.7 17l1.775-1.775l-.7-.7L7 16.3l-1.775-1.775l-.7.7L6.3 17l-1.775 1.775l.7.7Z"/>`),
		g.Group(children),
	)
}