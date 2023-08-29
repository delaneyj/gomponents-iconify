package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotTabletOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3ZM4 6H3v12h1V6Zm2 12h12V6H6v12ZM20 6v12h1V6h-1Zm0 0h1h-1ZM4 6H3h1Zm9 11h4v-4h-1.5v2.5H13V17Zm-6-6h1.5V8.5H11V7H7v4Z"/>`),
		g.Group(children),
	)
}