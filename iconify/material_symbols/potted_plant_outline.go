package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PottedPlantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.55 20h6.9l1-4h-8.9l1 4Zm0 2q-.7 0-1.225-.425t-.7-1.1L5.5 16h13l-1.125 4.475q-.175.675-.7 1.1T15.45 22h-6.9ZM5 14h14v-2H5v2Zm7-6q0-2.5 1.75-4.25T18 2q0 2.25-1.425 3.9T13 7.9V10h8v4q0 .825-.588 1.413T19 16H5q-.825 0-1.413-.588T3 14v-4h8V7.9q-2.15-.35-3.575-2T6 2q2.5 0 4.25 1.75T12 8Z"/>`),
		g.Group(children),
	)
}