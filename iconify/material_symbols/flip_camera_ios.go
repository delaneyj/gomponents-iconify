package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipCameraIos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L9 3h6l1.85 2H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21H4Zm8-3q1.975 0 3.4-1.338t1.55-3.312l.75.7L18.75 13l-2.5-2.5l-2.5 2.5l1.05 1.05l.65-.65q-.15 1.325-1.125 2.212T12 16.5q-.325 0-.638-.063t-.612-.187l-1.1 1.1q.55.3 1.137.475T12 18Zm-4.25-2.5l2.5-2.5l-1.05-1.05l-.65.65q.15-1.325 1.125-2.212T12 9.5q.325 0 .638.063t.612.187l1.1-1.1q-.55-.3-1.137-.475T12 8q-1.975 0-3.4 1.338T7.05 12.65l-.75-.7L5.25 13l2.5 2.5Z"/>`),
		g.Group(children),
	)
}