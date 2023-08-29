package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PottedPlantOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.55 20h6.9l1-4h-8.9l1 4Zm0 2q-.7 0-1.225-.425t-.7-1.1L5.5 16h13l-1.125 4.475q-.175.675-.7 1.1T15.45 22h-6.9ZM5 14h14v-2H5v2ZM6.75 2.05q2.225.275 3.738 1.962T12 8q0-2.3 1.513-3.988T17.25 2.05q.3-.05.525.175t.175.525q-.25 1.975-1.625 3.388T13 7.9V10h7q.425 0 .713.288T21 11v3q0 .825-.588 1.413T19 16H5q-.825 0-1.413-.588T3 14v-3q0-.425.288-.713T4 10h7V7.9q-1.95-.35-3.325-1.763T6.05 2.75q-.05-.3.175-.525t.525-.175Z"/>`),
		g.Group(children),
	)
}