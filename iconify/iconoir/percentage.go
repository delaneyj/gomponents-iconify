package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percentage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 19a2 2 0 1 1 0-4a2 2 0 0 1 0 4ZM7 9a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm12-4L5 19"/>`),
		g.Group(children),
	)
}