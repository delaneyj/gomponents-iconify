package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conversations(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.51 36.19l-.55-2A21.5 21.5 0 1 0 34.14 43l2 .55l7.2 1.93a1.67 1.67 0 0 0 2-1.18a1.61 1.61 0 0 0 0-.87Zm-10-14.39a2.2 2.2 0 1 1-2.2 2.2a2.23 2.23 0 0 1 2.22-2.2Zm-19.04 4.4a2.2 2.2 0 1 1 2.2-2.2a2.23 2.23 0 0 1-2.2 2.2Zm9.53 0a2.2 2.2 0 1 1 2.2-2.2a2.23 2.23 0 0 1-2.2 2.2Z"/>`),
		g.Group(children),
	)
}