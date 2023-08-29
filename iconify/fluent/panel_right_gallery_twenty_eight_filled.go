package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelRightGalleryTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.997 7.75A3.75 3.75 0 0 0 22.247 4H5.75A3.75 3.75 0 0 0 2 7.75v11.5A3.75 3.75 0 0 0 5.75 23h16.497a3.75 3.75 0 0 0 3.75-3.75V17.5H18v4H5.75a2.25 2.25 0 0 1-2.25-2.25V7.75A2.25 2.25 0 0 1 5.75 5.5H18v4h7.997V7.75ZM18 16h7.997v-5H18v5Z"/>`),
		g.Group(children),
	)
}