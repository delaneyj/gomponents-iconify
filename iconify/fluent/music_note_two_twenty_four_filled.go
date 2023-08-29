package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteTwoTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 2.75a.75.75 0 0 0-.965-.718l-10 3a.75.75 0 0 0-.535.718v9.877a3.5 3.5 0 1 0 1.496 2.702a.756.756 0 0 0 .004-.079v-7.942l8.5-2.55v5.87a3.5 3.5 0 1 0 1.496 2.702a.764.764 0 0 0 .004-.08V2.75Z"/>`),
		g.Group(children),
	)
}