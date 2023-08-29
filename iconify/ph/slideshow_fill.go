package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideshowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 64v128a16 16 0 0 1-16 16H64a16 16 0 0 1-16-16V64a16 16 0 0 1 16-16h128a16 16 0 0 1 16 16Zm24-16a8 8 0 0 0-8 8v144a8 8 0 0 0 16 0V56a8 8 0 0 0-8-8ZM24 48a8 8 0 0 0-8 8v144a8 8 0 0 0 16 0V56a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}