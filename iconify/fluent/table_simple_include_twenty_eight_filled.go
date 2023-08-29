package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleIncludeTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 6.75A3.75 3.75 0 0 1 6.75 3h6.5v9.955a3.77 3.77 0 0 0-.295.295H3v-6.5ZM22.25 12A3.74 3.74 0 0 1 25 13.2V6.75A3.75 3.75 0 0 0 21.25 3h-6.5v9.135a3.753 3.753 0 0 1 1-.135h6.5ZM12 15.75c0-.346.047-.682.135-1H3v6.5A3.75 3.75 0 0 0 6.75 25h6.45a3.74 3.74 0 0 1-1.2-2.75v-6.5Zm1 0A2.75 2.75 0 0 1 15.75 13h6.5A2.75 2.75 0 0 1 25 15.75v6.5A2.75 2.75 0 0 1 22.25 25h-6.5A2.75 2.75 0 0 1 13 22.25v-6.5Z"/>`),
		g.Group(children),
	)
}