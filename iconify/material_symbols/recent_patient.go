package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecentPatient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20v-2.8q0-.85.438-1.563T5.6 14.55q1.55-.775 3.15-1.163T12 13q.5 0 1 .038t1 .112V20H4Zm8-8q-1.65 0-2.825-1.175T8 8q0-1.65 1.175-2.825T12 4q1.65 0 2.825 1.175T16 8q0 1.65-1.175 2.825T12 12Zm6 12v-5h-2v-6h6l-2 4h2l-4 7Z"/>`),
		g.Group(children),
	)
}