package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsBrightnessRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.5 16l1.15 1.15q.15.15.35.15t.35-.15L13.5 16h2q.2 0 .35-.15t.15-.35v-2l1.15-1.15q.15-.15.15-.35t-.15-.35L16 10.5v-2q0-.2-.15-.35T15.5 8h-2l-1.15-1.15Q12.2 6.7 12 6.7t-.35.15L10.5 8h-2q-.2 0-.35.15T8 8.5v2l-1.15 1.15q-.15.15-.15.35t.15.35L8 13.5v2q0 .2.15.35t.35.15h2Zm1.5-1V9q1.25 0 2.125.875T15 12q0 1.25-.875 2.125T12 15Zm-8 5q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}