package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsBrightness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 17.5l1.5-1.5H16v-2.5l1.5-1.5l-1.5-1.5V8h-2.5L12 6.5L10.5 8H8v2.5L6.5 12L8 13.5V16h2.5l1.5 1.5Zm0-2.5V9q1.25 0 2.125.875T15 12q0 1.25-.875 2.125T12 15Zm-8 5q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}