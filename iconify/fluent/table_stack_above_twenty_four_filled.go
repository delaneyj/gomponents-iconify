package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackAboveTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.25 4.5a.75.75 0 0 0 0-1.5H3.75a.75.75 0 0 0 0 1.5h16.5Zm-5.75 5v5h-5v-5h5Zm-5 6.5h5v5h-5v-5ZM8 9.5v5H3v-4.25a.75.75 0 0 1 .75-.75H8ZM8 16H3v1.75A3.25 3.25 0 0 0 6.25 21H8v-5Zm13 0h-5v5h1.75A3.25 3.25 0 0 0 21 17.75V16Zm-5-1.5h5v-4.25a.75.75 0 0 0-.75-.75H16v5Z"/>`),
		g.Group(children),
	)
}