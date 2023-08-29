package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivacyScreenOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.45 23.3l-3.3-3.3H3.4l6.875-6.875l-1.4-1.4L2 18.6v-4.175l4.775-4.8l-1.4-1.4L2 11.6V6q0-.25.1-.438t.25-.362L.65 3.5l1.425-1.425l19.8 19.8L20.45 23.3ZM22 19.15l-8.875-8.875L19.4 4h.6q.825 0 1.413.588T22 6v13.15ZM11.725 8.875l-2.1-2.1L12.4 4h4.175l-4.85 4.875Zm-3.5-3.5L6.85 4H9.6L8.225 5.375Z"/>`),
		g.Group(children),
	)
}