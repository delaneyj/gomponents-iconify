package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraFrontOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.85 22.5l-1.4-1.4l1.1-1.1H5v-2h4.55l-1.1-1.1l1.4-1.4l3.5 3.5l-3.5 3.5ZM5 17V2h14v15h-2v-2.05q-1-.725-2.263-1.088T12 13.5q-1.475 0-2.738.363T7 14.95v1.975L6.925 17H5Zm2-3.85q1.2-.575 2.463-.863T12 12q1.275 0 2.538.288T17 13.15V4H7v9.15ZM14 20v-2h5v2h-5Zm-2-9q-1.25 0-2.125-.875T9 8q0-1.25.875-2.125T12 5q1.25 0 2.125.875T15 8q0 1.25-.875 2.125T12 11Zm0-2q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9Zm0 4.5ZM12 8Z"/>`),
		g.Group(children),
	)
}