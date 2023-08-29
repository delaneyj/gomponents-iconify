package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 19v2h-2v-2H4v2H2v-2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h9.81a6.481 6.481 0 0 1 4.69-2c1.843 0 3.508.767 4.69 2H22a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1Zm-5.5-5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm0-2a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5ZM4 13v2h2v-2H4Zm4 0v2h2v-2H8Z"/>`),
		g.Group(children),
	)
}