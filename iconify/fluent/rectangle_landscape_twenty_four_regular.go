package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleLandscapeTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 7.25A3.25 3.25 0 0 1 5.25 4h13.5A3.25 3.25 0 0 1 22 7.25v9.5A3.25 3.25 0 0 1 18.75 20H5.25A3.25 3.25 0 0 1 2 16.75v-9.5ZM5.25 5.5A1.75 1.75 0 0 0 3.5 7.25v9.5c0 .966.784 1.75 1.75 1.75h13.5a1.75 1.75 0 0 0 1.75-1.75v-9.5a1.75 1.75 0 0 0-1.75-1.75H5.25Z"/>`),
		g.Group(children),
	)
}