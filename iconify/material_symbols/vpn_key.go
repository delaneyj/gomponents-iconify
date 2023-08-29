package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VpnKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18q-2.5 0-4.25-1.75T1 12q0-2.5 1.75-4.25T7 6q2.025 0 3.538 1.137T12.65 10H23v4h-2v4h-4v-4h-4.35q-.6 1.725-2.113 2.863T7 18Zm0-4q.825 0 1.413-.588T9 12q0-.825-.588-1.413T7 10q-.825 0-1.413.588T5 12q0 .825.588 1.413T7 14Z"/>`),
		g.Group(children),
	)
}