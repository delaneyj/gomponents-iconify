package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 9V7.5H8A1.5 1.5 0 0 0 9.5 6v-.1a1.4 1.4 0 0 0-1.4-1.4h-.6A1.5 1.5 0 0 0 6 6m1 4.5h1"/>`),
		g.Group(children),
	)
}