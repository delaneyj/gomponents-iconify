package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConferenceRoomSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 7.998a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0ZM7.598 2.01A.5.5 0 0 0 7 2.5v10.997a.5.5 0 0 0 .598.49l5-1a.5.5 0 0 0 .402-.49V3.5a.5.5 0 0 0-.402-.49l-5-1ZM8 12.887V3.11l4 .8v8.177l-4 .8Zm-2 .11v-1H4V4h2V3H3.5a.5.5 0 0 0-.5.5v8.997a.5.5 0 0 0 .5.5H6Z"/>`),
		g.Group(children),
	)
}