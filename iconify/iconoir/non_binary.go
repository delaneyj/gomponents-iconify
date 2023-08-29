package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NonBinary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9a6 6 0 1 1 0 12a6 6 0 0 1 0-12Zm0 0V3M9 4l6 3m0-3L9 7"/>`),
		g.Group(children),
	)
}