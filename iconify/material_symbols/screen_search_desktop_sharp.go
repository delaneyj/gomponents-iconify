package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenSearchDesktopSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.45 15.05l1.1-1.05l-2.1-2.1q.275-.425.413-.9T15 10q0-1.475-1.038-2.488T11.5 6.5q-1.425 0-2.463 1.012T8 10q0 1.475 1.038 2.488T11.5 13.5q.525 0 .988-.138t.912-.412l2.05 2.1ZM11.5 12q-.825 0-1.413-.588T9.5 10q0-.825.588-1.413T11.5 8q.8 0 1.4.588T13.5 10q0 .825-.588 1.413T11.5 12ZM1 21v-2h22v2H1Zm1-3V3h20v15H2Z"/>`),
		g.Group(children),
	)
}