package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 5.75A2.75 2.75 0 0 1 4.75 3h10.5A2.75 2.75 0 0 1 18 5.75V10h-1V5.75A1.75 1.75 0 0 0 15.25 4H4.75A1.75 1.75 0 0 0 3 5.75v6.5c0 .966.784 1.75 1.75 1.75H9v1H4.75A2.75 2.75 0 0 1 2 12.25v-6.5ZM11.5 11a1.5 1.5 0 0 0-1.5 1.5v4a1.5 1.5 0 0 0 1.5 1.5h6a1.5 1.5 0 0 0 1.5-1.5v-4a1.5 1.5 0 0 0-1.5-1.5h-6Z"/>`),
		g.Group(children),
	)
}