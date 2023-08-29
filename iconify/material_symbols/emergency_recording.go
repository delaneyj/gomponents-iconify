package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyRecording(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17h2v-3.275l3 1.725l1-1.725L12 12l3-1.725l-1-1.725l-3 1.725V7H9v3.275L6 8.55l-1 1.725L8 12l-3 1.725l1 1.725l3-1.725V17Zm-5 3q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h12q.825 0 1.413.588T18 6v4.5l4-4v11l-4-4V18q0 .825-.588 1.413T16 20H4Z"/>`),
		g.Group(children),
	)
}