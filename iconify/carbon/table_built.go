package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBuilt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 8h-4V4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v18a2.002 2.002 0 0 0 2 2h4v4a2.002 2.002 0 0 0 2 2h18a2.002 2.002 0 0 0 2-2V10a2.002 2.002 0 0 0-2-2Zm-6 14h-8v-5h8Zm0-7h-8v-5h8Zm-10 0H4v-5h8ZM22 4v4H4V4ZM4 22v-5h8v5Zm24 6H10v-4h12a2.002 2.002 0 0 0 2-2V10h4Z"/>`),
		g.Group(children),
	)
}