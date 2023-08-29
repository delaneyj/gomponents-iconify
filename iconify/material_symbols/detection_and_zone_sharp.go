package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectionAndZoneSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 7V2h5v2H4v3H2Zm0 15v-5h2v3h3v2H2Zm15 0v-2h3v-3h2v5h-5Zm3-15V4h-3V2h5v5h-2Zm-6.5 2.5q-.825 0-1.413-.588T11.5 7.5q0-.825.588-1.413T13.5 5.5q.825 0 1.413.588T15.5 7.5q0 .825-.588 1.413T13.5 9.5ZM8.8 18l1-5.1l-1.8.7V17H6v-4.7l3.95-1.7q.875-.375 1.288-.487T12 10q.525 0 .975.275T13.7 11l1 1.6q.65 1.05 1.763 1.725T19 15v2q-1.65 0-3.088-.688T13.5 14.5l-.7 3.5h-4Z"/>`),
		g.Group(children),
	)
}