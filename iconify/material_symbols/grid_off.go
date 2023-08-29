package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 19.15l-2-2V16h-1.15l-2-2H20v-4h-4v3.15l-2-2V10h-1.15l-2-2H14V4h-4v3.15l-2-2V4H6.85l-2-2H20q.825 0 1.413.588T22 4v15.15ZM16 8h4V4h-4v4Zm4.5 15.3L19.15 22H4q-.825 0-1.413-.588T2 20V4.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM16 20h1.15L16 18.8V20Zm-6-6h1.15L10 12.8V14Zm0 6h4v-3.2l-.85-.8H10v4ZM4 8h1.15L4 6.8V8Zm0 6h4v-3.2l-.85-.8H4v4Zm4 6v-4H4v4h4Z"/>`),
		g.Group(children),
	)
}