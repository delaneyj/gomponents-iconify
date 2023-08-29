package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.236 7a3 3 0 1 0-4.472-4a3 3 0 0 0 4.472 4Zm0 0L20 18M7.236 17a3 3 0 1 1-4.472 4a3 3 0 0 1 4.472-4Zm0 0L20 6"/>`),
		g.Group(children),
	)
}