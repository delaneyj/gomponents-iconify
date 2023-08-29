package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLandscapeDataTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.75A2.75 2.75 0 0 1 4.75 4h14.5A2.75 2.75 0 0 1 22 6.75v10.5A2.75 2.75 0 0 1 19.25 20H4.75A2.75 2.75 0 0 1 2 17.25V6.75ZM12 7a2 2 0 0 0-2 2v6a2 2 0 1 0 4 0V9a2 2 0 0 0-2-2Zm-5 5a2 2 0 0 0-2 2v1a2 2 0 1 0 4 0v-1a2 2 0 0 0-2-2Zm8 0v3a2 2 0 1 0 4 0v-3a2 2 0 1 0-4 0Z"/>`),
		g.Group(children),
	)
}