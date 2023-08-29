package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fishing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 0v10c0 6-10 6-10 0v-4l2 2"/>`),
		g.Group(children),
	)
}