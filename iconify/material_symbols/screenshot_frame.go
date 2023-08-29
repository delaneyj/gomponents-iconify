package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotFrame(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7V4q0-.825.588-1.413T7 2h3v2H7v3H5Zm2 15q-.825 0-1.413-.588T5 20v-3h2v3h3v2H7ZM17 7V4h-3V2h3q.825 0 1.413.588T19 4v3h-2Zm-3 15v-2h3v-3h2v3q0 .825-.588 1.413T17 22h-3Z"/>`),
		g.Group(children),
	)
}