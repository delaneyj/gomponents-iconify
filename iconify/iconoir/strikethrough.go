package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h18m-4.714-9h-6.218a4.068 4.068 0 0 0-1.286 7.927L12 12m-6 9h7.932a4.068 4.068 0 0 0 3.58-6"/>`),
		g.Group(children),
	)
}