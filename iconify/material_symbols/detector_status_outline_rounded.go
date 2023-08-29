package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorStatusOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.95 18.175l3.525-3.525q.3-.3.713-.313t.712.288q.3.3.3.713t-.3.712l-4.25 4.25q-.3.3-.7.3t-.7-.3L8.1 18.15q-.3-.3-.287-.7t.312-.7q.3-.275.7-.287t.7.287l1.425 1.425ZM5 5v1h14V5H5Zm3.1 3l.3 1h7.2l.3-1H8.1Zm.3 3q-.65 0-1.175-.388T6.5 9.6L6 8H5q-.825 0-1.413-.587T3 6V3h18v3q0 .825-.588 1.413T19 8h-1l-.65 1.7q-.225.575-.725.938T15.5 11H8.4ZM5 5v1v-1Z"/>`),
		g.Group(children),
	)
}