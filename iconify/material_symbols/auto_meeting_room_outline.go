package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoMeetingRoomOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21v-2h2V3h10v1h4v8h-2V6h-2v15H1Zm8-8q.425 0 .713-.288T10 12q0-.425-.288-.713T9 11q-.425 0-.713.288T8 12q0 .425.288.713T9 13Zm-4 6h6V5H5v14Zm14 3l1.25-2.75L23 18l-2.75-1.25L19 14l-1.25 2.75L15 18l2.75 1.25L19 22ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}