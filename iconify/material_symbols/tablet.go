package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3Zm3-2h12V6H6v12Z"/>`),
		g.Group(children),
	)
}