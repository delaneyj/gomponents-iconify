package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoardSplitTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3A3.75 3.75 0 0 0 3 6.75V13h13V3H6.75ZM16 14.5H3v6.75A3.75 3.75 0 0 0 6.75 25H16V14.5ZM21.25 25H17.5v-7H25v3.25A3.75 3.75 0 0 1 21.25 25ZM25 16.5h-7.5v-5H25v5ZM17.5 3v7H25V6.75A3.75 3.75 0 0 0 21.25 3H17.5Z"/>`),
		g.Group(children),
	)
}