package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiningTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 .5L16 8l7.5 4l-7.5 4l-4 7.5L8 16L.5 12L8 8l4-7.5Z"/>`),
		g.Group(children),
	)
}