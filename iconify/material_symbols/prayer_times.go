package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrayerTimes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q2.075 0 3.538-1.45T17 12q0-.2-.013-.4t-.062-.4q-.275 1.175-1.225 1.938t-2.225.762q-1.5 0-2.525-1.025T9.925 10.35q0-1.15.65-2.063T12.275 7H12Q9.9 7 8.45 8.463T7 12q0 2.1 1.45 3.55T12 17Zm1.025-5.5l1.475-1.075l1.45 1.075l-.575-1.7l1.475-1.075l-1.8.025L14.5 7l-.55 1.75l-1.825-.025L13.6 9.8l-.575 1.7ZM12 23.3L8.65 20H4v-4.65L.7 12L4 8.65V4h4.65L12 .7L15.35 4H20v4.65L23.3 12L20 15.35V20h-4.65L12 23.3Z"/>`),
		g.Group(children),
	)
}