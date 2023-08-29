package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Category(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 22.141V18a2 2 0 0 0-2-2h-8v-4h2a2.002 2.002 0 0 0 2-2V4a2.002 2.002 0 0 0-2-2h-6a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h2v4H7a2 2 0 0 0-2 2v4.142a4 4 0 1 0 2 0V18h8v4.142a4 4 0 1 0 2 0V18h8v4.141a4 4 0 1 0 2 0ZM13 4h6l.001 6H13ZM8 26a2 2 0 1 1-2-2a2.002 2.002 0 0 1 2 2Zm10 0a2 2 0 1 1-2-2a2.003 2.003 0 0 1 2 2Zm8 2a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}