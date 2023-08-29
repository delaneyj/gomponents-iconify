package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundedMirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 10v4a8 8 0 1 1-16 0v-4a8 8 0 1 1 16 0Zm-2.5-5.5L13 8m6-1l-7.5 6"/>`),
		g.Group(children),
	)
}