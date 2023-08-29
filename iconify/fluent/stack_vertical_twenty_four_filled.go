package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackVerticalTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.25A3.25 3.25 0 0 1 6.25 3h11.5A3.25 3.25 0 0 1 21 6.25V20a1 1 0 1 1-2 0v-3.75c0-.69-.56-1.25-1.25-1.25H6.25C5.56 15 5 15.56 5 16.25V20a1 1 0 1 1-2 0V6.25Zm2 7A3.24 3.24 0 0 1 6.25 13h11.5c.443 0 .865.089 1.25.25v-2c0-.69-.56-1.25-1.25-1.25H6.25C5.56 10 5 10.56 5 11.25v2Zm14-7C19 5.56 18.44 5 17.75 5H6.25C5.56 5 5 5.56 5 6.25v2A3.24 3.24 0 0 1 6.25 8h11.5c.443 0 .865.089 1.25.25v-2Z"/>`),
		g.Group(children),
	)
}