package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.25A3.25 3.25 0 0 1 5.25 3h13.5A3.25 3.25 0 0 1 22 6.25V12h-1.5V6.25a1.75 1.75 0 0 0-1.75-1.75H5.25A1.75 1.75 0 0 0 3.5 6.25v9.5c0 .966.784 1.75 1.75 1.75H11V19H5.25A3.25 3.25 0 0 1 2 15.75v-9.5ZM14 13a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2h-7Z"/>`),
		g.Group(children),
	)
}