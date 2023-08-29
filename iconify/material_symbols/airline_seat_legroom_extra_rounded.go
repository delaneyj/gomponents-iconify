package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomExtraRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17H4q-.825 0-1.413-.588T2 15V4q0-.425.288-.713T3 3q.425 0 .713.288T4 4v11h8q.425 0 .713.288T13 16q0 .425-.288.713T12 17Zm3-3H8.5q-1.25 0-2.125-.875T5.5 11V3h6v6h3q.575 0 1.05.313t.75.837l3.4 6.95l1.1-.5q.575-.275 1.163-.088t.887.738q.3.575.088 1.175t-.788.875l-2.875 1.3q-.375.175-.75.025t-.55-.5L15 14Z"/>`),
		g.Group(children),
	)
}