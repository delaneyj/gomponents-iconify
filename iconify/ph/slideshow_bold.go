package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SlideshowBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 44H72a20 20 0 0 0-20 20v128a20 20 0 0 0 20 20h112a20 20 0 0 0 20-20V64a20 20 0 0 0-20-20Zm-4 144H76V68h104Zm64-132v144a12 12 0 0 1-24 0V56a12 12 0 0 1 24 0ZM36 56v144a12 12 0 0 1-24 0V56a12 12 0 0 1 24 0Z"/>`),
		g.Group(children),
	)
}