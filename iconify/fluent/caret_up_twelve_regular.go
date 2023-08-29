package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUpTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.076 7.383A1 1 0 0 0 4 8h4a1 1 0 0 0 .707-1.707l-2-2a1 1 0 0 0-1.414 0l-2 2a1 1 0 0 0-.217 1.09ZM4 7l1.823-1.823a.25.25 0 0 1 .354 0L8 7H4Z"/>`),
		g.Group(children),
	)
}