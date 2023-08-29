package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 29a1 1 0 1 1 0-2h24a1 1 0 1 1 0 2H4Zm1-7.5A3.5 3.5 0 0 0 8.5 25h3a3.5 3.5 0 0 0 3.5-3.5v-15A3.5 3.5 0 0 0 11.5 3h-3A3.5 3.5 0 0 0 5 6.5v15ZM8.5 23A1.5 1.5 0 0 1 7 21.5v-15A1.5 1.5 0 0 1 8.5 5h3A1.5 1.5 0 0 1 13 6.5v15a1.5 1.5 0 0 1-1.5 1.5h-3Zm8.5-1.5a3.5 3.5 0 0 0 3.5 3.5h3a3.5 3.5 0 0 0 3.5-3.5v-9A3.5 3.5 0 0 0 23.5 9h-3a3.5 3.5 0 0 0-3.5 3.5v9Zm3.5 1.5a1.5 1.5 0 0 1-1.5-1.5v-9a1.5 1.5 0 0 1 1.5-1.5h3a1.5 1.5 0 0 1 1.5 1.5v9a1.5 1.5 0 0 1-1.5 1.5h-3Z"/>`),
		g.Group(children),
	)
}