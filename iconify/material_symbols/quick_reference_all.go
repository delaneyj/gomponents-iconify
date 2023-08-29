package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickReferenceAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 19q1.05 0 1.775-.725T19 16.5q0-1.05-.725-1.775T16.5 14q-1.05 0-1.775.725T14 16.5q0 1.05.725 1.775T16.5 19Zm5.1 4l-2.7-2.675q-.525.325-1.137.5T16.5 21q-1.875 0-3.187-1.312T12 16.5q0-1.875 1.313-3.188T16.5 12q1.875 0 3.188 1.313T21 16.5q0 .675-.188 1.288t-.512 1.137L23 21.6L21.6 23ZM5 22q-.825 0-1.413-.588T3 20V4q0-.825.588-1.413T5 2h8l6 6v2.5q-.6-.25-1.225-.375T16.5 10q-1.4 0-2.613.525T11.8 12H7v2h3.5q-.2.475-.313.975T10.025 16H7v2h3.175q.275 1.25 1 2.288T13.025 22H5Zm7-13h5l-5-5l5 5l-5-5v5Z"/>`),
		g.Group(children),
	)
}