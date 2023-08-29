package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTypingAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7 7.5h1m-4 0h1m5 0h1m3.5 7h-7a7 7 0 1 1 7-7v7Z"/>`),
		g.Group(children),
	)
}