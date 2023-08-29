package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextParagraphSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.25 2a.75.75 0 0 0-.75.75v2a.75.75 0 0 1-.75.75H9.56l.72-.72a.75.75 0 1 0-1.06-1.06l-2 2a.75.75 0 0 0 0 1.06l2 2a.75.75 0 1 0 1.06-1.06L9.56 7h2.19A2.25 2.25 0 0 0 14 4.75v-2a.75.75 0 0 0-.75-.75ZM5.72 8.22a.75.75 0 0 1 1.06 0l2 2a.75.75 0 0 1 0 1.06l-2 2a.75.75 0 0 1-1.06-1.06l.72-.72H1.75a.75.75 0 0 1 0-1.5h4.69l-.72-.72a.75.75 0 0 1 0-1.06Z"/>`),
		g.Group(children),
	)
}