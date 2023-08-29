package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 23q-.825 0-1.413-.588T4 21V10q0-.825.588-1.413T6 8h3v2H6v11h12V10h-3V8h3q.825 0 1.413.588T20 10v11q0 .825-.588 1.413T18 23H6Zm5-7V4.825l-1.6 1.6L8 5l4-4l4 4l-1.4 1.425l-1.6-1.6V16h-2Z"/>`),
		g.Group(children),
	)
}