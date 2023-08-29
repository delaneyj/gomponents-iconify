package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowReplyTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m3.707 9l3.39 3.39a.5.5 0 0 1-.639.764l-.069-.057l-4.243-4.243a.5.5 0 0 1-.057-.638l.057-.07L6.39 3.905a.5.5 0 0 1 .765.638l-.058.07L3.707 8H10a7.5 7.5 0 0 1 7.496 7.258l.004.243a.5.5 0 0 1-1 0a6.5 6.5 0 0 0-6.267-6.496L10 9H3.707l3.39 3.39L3.706 9Z"/>`),
		g.Group(children),
	)
}