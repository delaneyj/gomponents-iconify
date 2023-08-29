package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelicopterOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 13V7Q6.5 7 4.75 8.75T3 13h6Zm4 4V7h-2v8H3v2h10Zm2-3.2l6-.6V12h-6v1.8ZM13 22H3v-2h10v2Zm2-3H3q-.825 0-1.413-.588T1 17v-4q0-3.35 2.325-5.675T9 5h6v5h5l1-2h2v7l-8 .8V19Zm4-15H3V2h16v2Zm-4 9.8V12v1.8ZM13 17Z"/>`),
		g.Group(children),
	)
}