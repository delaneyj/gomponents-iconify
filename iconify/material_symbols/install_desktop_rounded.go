package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstallDesktopRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20v-1H4q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h7q.425 0 .713.288T12 4q0 .425-.288.713T11 5H4v12h16v-2q0-.425.288-.713T21 14q.425 0 .713.288T22 15v2q0 .825-.588 1.413T20 19h-4v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20Zm8-9.825V4q0-.425.288-.713T17 3q.425 0 .713.288T18 4v6.175L19.9 8.3q.275-.275.688-.288t.712.288q.275.275.275.7t-.275.7l-3.6 3.6q-.3.3-.7.3t-.7-.3l-3.6-3.6q-.275-.275-.287-.687T12.7 8.3q.275-.275.7-.275t.7.275l1.9 1.875Z"/>`),
		g.Group(children),
	)
}