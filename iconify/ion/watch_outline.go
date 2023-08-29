package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<rect width="288" height="288" x="112" y="112" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" rx="64" ry="64"/><path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M176 112V40a8 8 0 0 1 8-8h144a8 8 0 0 1 8 8v72m0 288v72a8 8 0 0 1-8 8H184a8 8 0 0 1-8-8v-72"/>`),
		g.Group(children),
	)
}