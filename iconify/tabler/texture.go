package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Texture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3L3 6m18 12l-3 3M11 3l-8 8m13-8L3 16M21 3L3 21M21 8L8 21m13-8l-8 8"/>`),
		g.Group(children),
	)
}