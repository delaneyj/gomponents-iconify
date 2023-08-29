package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vignette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm8-4q2.45 0 4.225-1.188T18 12q0-1.625-1.775-2.813T12 8Q9.55 8 7.775 9.188T6 12q0 1.625 1.775 2.813T12 16Z"/>`),
		g.Group(children),
	)
}