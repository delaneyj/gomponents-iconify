package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLandscapeSplitTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6.5V4ZM12 20h8a2 2 0 0 0 2-2v-6h-6a2 2 0 0 1-2-2V4h-2v16Zm4-9.5h5.5l-6-6V10a.5.5 0 0 0 .5.5Z"/>`),
		g.Group(children),
	)
}