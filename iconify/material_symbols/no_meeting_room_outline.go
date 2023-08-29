package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoMeetingRoomOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 16.15l-2-2V6h-2v6.15l-2-2V5H7.85l-2-2H15v1h4v12.15Zm.8 6.45L15 17.8V21H3v-2h2V7.8L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM7 19h6v-3.2l-6-6V19Zm3.425-11.425ZM10 12.8Z"/>`),
		g.Group(children),
	)
}