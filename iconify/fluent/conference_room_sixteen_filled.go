package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConferenceRoomSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.183 2.113a.5.5 0 0 1 .415-.103l5 1A.5.5 0 0 1 13 3.5v8.997a.5.5 0 0 1-.402.49l-5 1a.5.5 0 0 1-.598-.49V2.5m3 5.498a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0ZM6 3H3.5a.5.5 0 0 0-.5.5v8.997a.5.5 0 0 0 .5.5H6V3Zm1.183-.887A.5.5 0 0 0 7 2.5Z"/>`),
		g.Group(children),
	)
}