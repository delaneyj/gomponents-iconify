package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeveloperMode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.6 16.6L4 12l4.6-4.6L10 8.85L6.85 12L10 15.15L8.6 16.6ZM5 17h2v1h10v-1h2v4q0 .825-.588 1.413T17 23H7q-.825 0-1.413-.588T5 21v-4ZM7 7H5V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v4h-2V6H7v1Zm8.4 9.6L14 15.15L17.15 12L14 8.85l1.4-1.45L20 12l-4.6 4.6Z"/>`),
		g.Group(children),
	)
}