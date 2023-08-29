package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForwardMiniLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.032 12L9 9.86v4.28L12.032 12ZM7.5 17.535a.5.5 0 0 1-.5-.5V6.965a.5.5 0 0 1 .788-.409l7.133 5.035a.499.499 0 0 1 0 .818l-7.133 5.034a.5.5 0 0 1-.288.092ZM16 7a1 1 0 1 1 2 0v10a1 1 0 1 1-2 0V7Z"/>`),
		g.Group(children),
	)
}