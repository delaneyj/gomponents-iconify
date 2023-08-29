package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZeroThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5Z"/>`),
		g.Group(children),
	)
}