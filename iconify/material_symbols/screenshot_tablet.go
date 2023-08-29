package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotTablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3Zm3-2h12V6H6v12Zm7-1h4v-4h-1.5v2.5H13V17Zm-6-6h1.5V8.5H11V7H7v4Z"/>`),
		g.Group(children),
	)
}