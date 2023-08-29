package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glasses(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 14a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm0 0V6m20 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm0 0V6m-8 8h-4"/>`),
		g.Group(children),
	)
}