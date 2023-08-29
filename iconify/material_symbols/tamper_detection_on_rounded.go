package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TamperDetectionOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.425 20q-.425 0-.788-.163T3 19.4L.675 17.075Q.4 16.8.4 16.388t.275-.688q.3-.3.713-.3t.687.3l.9.9V9.25q0-.325.225-.537t.525-.213q.325 0 .537.213t.213.537V13h1V7.75q0-.325.225-.537T6.225 7q.325 0 .538.213t.212.537V13h1V8.75q0-.325.225-.537T8.725 8q.325 0 .538.213t.212.537V13h1V9.75q0-.325.225-.537T11.225 9q.325 0 .538.213t.212.537V18q0 .825-.575 1.413T9.975 20h-5.55Zm9.55-3V8q0-.825-.587-1.413T11.975 6h-10V3q0-.825.588-1.413T3.974 1h12q.825 0 1.413.588T17.974 3v4.5l3.15-3.15q.25-.25.55-.125t.3.475v8.575q0 .35-.3.475t-.55-.125l-3.15-3.125V15q0 .825-.588 1.413t-1.41.587h-2Z"/>`),
		g.Group(children),
	)
}