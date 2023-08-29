package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleCheckmarkTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3h6.5v10.25H3v-6.5A3.75 3.75 0 0 1 6.75 3ZM3 14.75h10.25V25h-6.5A3.75 3.75 0 0 1 3 21.25v-6.5Zm11.75 0V25h6.5A3.75 3.75 0 0 0 25 21.25v-6.5H14.75Zm0-11.75v10.25H25v-6.5A3.75 3.75 0 0 0 21.25 3h-6.5Zm8.028 15.78l-3 3a.75.75 0 0 1-1.06 0l-1.498-1.498a.75.75 0 0 1 1.06-1.06l.968.967l2.47-2.47a.75.75 0 0 1 1.06 1.061Z"/>`),
		g.Group(children),
	)
}