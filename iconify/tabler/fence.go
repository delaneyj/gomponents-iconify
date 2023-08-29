package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fence(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12v4h16v-4zm2 4v4h4v-4m0-4V6L8 4L6 6v6m8 4v4h4v-4m0-4V6l-2-2l-2 2v6"/>`),
		g.Group(children),
	)
}