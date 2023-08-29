package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mention(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12.002V13a2 2 0 1 0 4 0v-1a7 7 0 1 0-4.406 6.502m.406-6.5a3 3 0 1 1 0-.004m0 .004v-.004m0 0V9"/>`),
		g.Group(children),
	)
}