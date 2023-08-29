package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VehicleCableCarTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.434 3.004a.5.5 0 1 1 .132.992L12 4.738V6a2.99 2.99 0 0 1-.764 2H13a3 3 0 0 1 3 3v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4a3 3 0 0 1 3-3h2a2 2 0 0 0 2-2V4.871L2.566 5.996a.5.5 0 1 1-.132-.992L11 3.862V3.5a.5.5 0 1 1 1 0v.229l5.434-.725ZM15 11a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v2h10v-2ZM5 14v1a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1H5Z"/>`),
		g.Group(children),
	)
}