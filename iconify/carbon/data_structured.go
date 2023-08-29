package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataStructured(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="9" cy="7" r="1" fill="currentColor"/><path fill="currentColor" d="M27 22.14V18a2 2 0 0 0-2-2h-8v-4h9a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h9v4H7a2 2 0 0 0-2 2v4.14a4 4 0 1 0 2 0V18h8v4h-3v8h8v-8h-3v-4h8v4.14a4 4 0 1 0 2 0ZM8 26a2 2 0 1 1-2-2a2 2 0 0 1 2 2Zm10-2v4h-4v-4ZM6 10V4h20v6Zm20 18a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}