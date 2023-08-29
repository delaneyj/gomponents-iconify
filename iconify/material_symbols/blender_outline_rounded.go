package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlenderOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22q-.825 0-1.413-.588T6 20v-1q0-1.175.513-2.175T7.85 15.15L7.25 11H5q-.825 0-1.413-.588T3 9V5q0-.825.588-1.413T5 3h5q0-.425.288-.713T11 2h2q.425 0 .713.288T14 3h2.825q.45 0 .763.35t.237.8l-1.675 11q.825.675 1.338 1.675T18 19v1q0 .825-.588 1.413T16 22H8ZM6.95 9L6.3 5H5v4h1.95ZM12 19q.425 0 .713-.288T13 18q0-.425-.288-.713T12 17q-.425 0-.713.288T11 18q0 .425.288.713T12 19Zm-2.3-5h4.6l1.35-9h-7.3l1.35 9ZM8 20h8v-1q0-1.25-.875-2.125T13 16h-2q-1.25 0-2.125.875T8 19v1Zm4-4Z"/>`),
		g.Group(children),
	)
}