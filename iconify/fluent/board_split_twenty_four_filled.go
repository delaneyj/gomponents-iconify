package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoardSplitTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 3A3.25 3.25 0 0 0 3 6.25V11h11V3H6.25ZM3 17.75V12.5h11V21H6.25A3.25 3.25 0 0 1 3 17.75ZM15.5 16v5h2.25A3.25 3.25 0 0 0 21 17.75V16h-5.5Zm5.5-1.5v-5h-5.5v5H21ZM21 8h-5.5V3h2.25A3.25 3.25 0 0 1 21 6.25V8Z"/>`),
		g.Group(children),
	)
}