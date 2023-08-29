package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.75 8A4.75 4.75 0 0 0 4 12.75v18.5A4.75 4.75 0 0 0 8.75 36h30.5A4.75 4.75 0 0 0 44 31.25v-18.5A4.75 4.75 0 0 0 39.25 8H8.75ZM6.5 12.75a2.25 2.25 0 0 1 2.25-2.25h30.5a2.25 2.25 0 0 1 2.25 2.25v18.5a2.25 2.25 0 0 1-2.25 2.25H8.75a2.25 2.25 0 0 1-2.25-2.25v-18.5Zm4.75 26.75a1.25 1.25 0 1 0 0 2.5h25.5a1.25 1.25 0 1 0 0-2.5h-25.5Z"/>`),
		g.Group(children),
	)
}