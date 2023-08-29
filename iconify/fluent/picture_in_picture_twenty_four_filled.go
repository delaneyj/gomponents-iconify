package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.25 3A3.25 3.25 0 0 0 2 6.25v9.5A3.25 3.25 0 0 0 5.25 19H11v-4a3 3 0 0 1 3-3h7c.35 0 .687.06 1 .17V6.25A3.25 3.25 0 0 0 18.75 3H5.25ZM22 13.268A1.99 1.99 0 0 0 21 13h-7a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-5a2 2 0 0 0-1-1.732Z"/>`),
		g.Group(children),
	)
}