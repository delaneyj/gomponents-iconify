package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConferenceRoomTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11.894 2.014l11.5 2.25A.75.75 0 0 1 24 5v18a.75.75 0 0 1-.606.736l-11.5 2.25A.75.75 0 0 1 11 25.25V2.75a.75.75 0 0 1 .894-.736ZM15 13a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-5-8.5v19H4.75a.75.75 0 0 1-.743-.648L4 22.75V5.25a.75.75 0 0 1 .648-.743L4.75 4.5H10Z"/>`),
		g.Group(children),
	)
}