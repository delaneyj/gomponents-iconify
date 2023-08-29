package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VillaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V8l13-5v9h1q0-.825.588-1.413T19 10q.825 0 1.413.588T21 12v9H3Zm2-2h4v-7h5V5.9L5 9.375V19Zm6 0h3v-3h2v3h3v-5h-8v5Zm-4-9Zm8 9Zm0-.5Z"/>`),
		g.Group(children),
	)
}