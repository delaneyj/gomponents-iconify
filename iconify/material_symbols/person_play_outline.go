package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonPlayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 12.4L1.6 8.5l3.9-3.9l3.9 3.9l-3.9 3.9ZM9 22v-5q-1.525-.125-3.025-.363T3 16l.5-2q2.1.575 4.213.788T12 15q2.175 0 4.288-.213T20.5 14l.5 2q-1.475.4-2.975.638T15 17v5H9ZM5.5 9.6l1.1-1.1l-1.1-1.1l-1.1 1.1l1.1 1.1ZM12 7q-1.25 0-2.125-.875T9 4q0-1.25.875-2.125T12 1q1.25 0 2.125.875T15 4q0 1.25-.875 2.125T12 7Zm0 7q-.825 0-1.413-.587T10 12q0-.825.588-1.413T12 10q.825 0 1.413.588T14 12q0 .825-.588 1.413T12 14Zm0-9q.425 0 .713-.287T13 4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4q0 .425.288.713T12 5Zm5.05 7l-1.7-3l1.7-3h3.4l1.7 3l-1.7 3h-3.4Zm1.15-2h1.1l.55-1l-.55-1h-1.1l-.55 1l.55 1ZM5.5 8.5ZM12 4Zm6.75 5Z"/>`),
		g.Group(children),
	)
}