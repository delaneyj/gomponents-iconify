package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MeetingRoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h2V3h10v1h4v15h2v2h-4V6h-2v15H3Zm8-8q.425 0 .713-.288T12 12q0-.425-.288-.713T11 11q-.425 0-.713.288T10 12q0 .425.288.713T11 13Z"/>`),
		g.Group(children),
	)
}