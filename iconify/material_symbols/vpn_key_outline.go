package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VpnKeyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18q-2.5 0-4.25-1.75T1 12q0-2.5 1.75-4.25T7 6q1.65 0 3.025.825T12.2 9H23v6h-2v3h-6v-3h-2.8q-.8 1.35-2.175 2.175T7 18Zm0-2q1.65 0 2.65-1.012T10.85 13H17v3h2v-3h2v-2H10.85q-.2-.975-1.2-1.988T7 8Q5.35 8 4.175 9.175T3 12q0 1.65 1.175 2.825T7 16Zm0-2q.825 0 1.413-.588T9 12q0-.825-.588-1.413T7 10q-.825 0-1.413.588T5 12q0 .825.588 1.413T7 14Zm0-2Z"/>`),
		g.Group(children),
	)
}