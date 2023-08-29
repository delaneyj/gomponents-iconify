package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scene(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22V7q0-.825-.588-1.413T18 5h-1v1.3q0 .3-.2.5t-.5.2h-5.6q-.35 0-.562-.375T10.1 5.9L12 1.8q.175-.375.513-.588T13.3 1h2.3q.6 0 1 .45T17 2.5V3h1q1.65 0 2.825 1.175T22 7v15h-2ZM5 22q-1.275 0-2.138-.863T2 19v-2.5q0-.625.438-1.063T3.5 15q.625 0 1.063.438T5 16.5V19h10v-2.5q0-.625.438-1.063T16.5 15q.625 0 1.063.438T18 16.5V19q0 1.275-.863 2.138T15 22H5Zm1-4v-1.5q0-.8-.525-1.525T4 14.05V12q0-.825.587-1.413T6 10h8q.825 0 1.413.588T16 12v2.05q-.95.2-1.475.925T14 16.5V18H6Z"/>`),
		g.Group(children),
	)
}