package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextParagraphTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.25 2a.75.75 0 0 1 .75.75v4a2.75 2.75 0 0 1-2.75 2.75h-5.69l2.72 2.72a.75.75 0 1 1-1.06 1.06l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 1 1 1.06 1.06L13.56 8h5.69c.69 0 1.25-.56 1.25-1.25v-4a.75.75 0 0 1 .75-.75ZM7.72 11.22a.75.75 0 0 0 0 1.06L10.44 15H2.75a.75.75 0 0 0 0 1.5h7.69l-2.72 2.72a.75.75 0 1 0 1.06 1.06l4-4a.75.75 0 0 0 0-1.06l-4-4a.75.75 0 0 0-1.06 0Z"/>`),
		g.Group(children),
	)
}