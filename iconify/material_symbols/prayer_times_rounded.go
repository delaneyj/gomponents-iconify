package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrayerTimesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q2.075 0 3.538-1.45T17 12q0-.2-.013-.4t-.062-.4q-.275 1.175-1.225 1.938t-2.225.762q-1.5 0-2.525-1.025T9.925 10.35q0-1.15.65-2.063T12.275 7H12Q9.9 7 8.45 8.463T7 12q0 2.1 1.45 3.55T12 17Zm1.025-5.5l1.475-1.075l1.45 1.075l-.575-1.7l1.475-1.075l-1.8.025L14.5 7l-.55 1.75l-1.825-.025L13.6 9.8l-.575 1.7ZM8.65 20H6q-.825 0-1.413-.588T4 18v-2.65l-1.9-1.925Q1.5 12.85 1.5 12t.6-1.425L4 8.65V6q0-.825.588-1.413T6 4h2.65l1.925-1.9q.575-.6 1.425-.6t1.425.6L15.35 4H18q.825 0 1.413.588T20 6v2.65l1.9 1.925q.6.575.6 1.425t-.6 1.425L20 15.35V18q0 .825-.588 1.413T18 20h-2.65l-1.925 1.9q-.575.6-1.425.6t-1.425-.6L8.65 20Z"/>`),
		g.Group(children),
	)
}