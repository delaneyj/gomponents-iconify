package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Publish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 20v-8.15l-2.6 2.6L7 13l5-5l5 5l-1.4 1.45l-2.6-2.6V20h-2ZM4 9V6q0-.825.588-1.413T6 4h12q.825 0 1.413.588T20 6v3h-2V6H6v3H4Z"/>`),
		g.Group(children),
	)
}