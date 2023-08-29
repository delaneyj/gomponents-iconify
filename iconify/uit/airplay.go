package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 3h-15A2.502 2.502 0 0 0 2 5.5v10A2.502 2.502 0 0 0 4.5 18h1a.5.5 0 0 0 0-1h-1A1.5 1.5 0 0 1 3 15.5v-10A1.5 1.5 0 0 1 4.5 4h15A1.5 1.5 0 0 1 21 5.5v10a1.5 1.5 0 0 1-1.5 1.5h-1a.5.5 0 0 0 0 1h1a2.502 2.502 0 0 0 2.5-2.5v-10A2.502 2.502 0 0 0 19.5 3zm-6.259 11.44a1.557 1.557 0 0 0-2.482-.002l-2.863 4.22A1.5 1.5 0 0 0 9.136 21h5.727a1.5 1.5 0 0 0 1.241-2.342l-2.863-4.219zM14.863 20H9.137a.5.5 0 0 1-.413-.781L11.587 15a.494.494 0 0 1 .413-.219a.492.492 0 0 1 .413.22l2.863 4.218a.5.5 0 0 1-.413.781z"/>`),
		g.Group(children),
	)
}