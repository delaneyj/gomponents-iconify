package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudAlerting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 17v5H4V6h11V4H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h8v4H8v2h16v-2h-4v-4h8a2 2 0 0 0 2-2v-5ZM18 28h-4v-4h4Z"/><path fill="currentColor" d="M29 14H17a1 1 0 0 1-.857-1.514l6-10a1 1 0 0 1 1.715 0l6 10A1 1 0 0 1 29 14Zm-10.234-2h8.468L23 4.944Z"/>`),
		g.Group(children),
	)
}