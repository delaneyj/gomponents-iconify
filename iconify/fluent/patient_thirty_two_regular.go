package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatientThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 14a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-2a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-6 12a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2H10ZM5 5.5A3.5 3.5 0 0 1 8.5 2h15A3.5 3.5 0 0 1 27 5.5v21a3.5 3.5 0 0 1-3.5 3.5h-15A3.5 3.5 0 0 1 5 26.5v-21ZM8.5 4A1.5 1.5 0 0 0 7 5.5V20h3v-3a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v3h3V5.5A1.5 1.5 0 0 0 23.5 4h-15ZM7 26.5A1.5 1.5 0 0 0 8.5 28h15a1.5 1.5 0 0 0 1.5-1.5V22H7v4.5ZM20 17h-8v3h8v-3Z"/>`),
		g.Group(children),
	)
}