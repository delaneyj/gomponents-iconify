package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloseCornerArrowTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M170.7 0v128L42.7 0L0 42.7l128 128H0l64 64l170.8.1l-.1-170.8l-64-64zM512 341.3l-64-64l-170.8-.1l.1 170.8l64 64V384l128 128l42.7-42.7l-128-128h128z"/>`),
		g.Group(children),
	)
}