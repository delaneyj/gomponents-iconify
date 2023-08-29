package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommunityFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19h3v-6.058L8 9.454l-4 3.488V19h3v-4h2v4Zm12 2H3a1 1 0 0 1-1-1v-7.513a1 1 0 0 1 .343-.754L6 8.544V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1Zm-5-10v2h2v-2h-2Zm0 4v2h2v-2h-2Zm0-8v2h2V7h-2Zm-4 0v2h2V7h-2Z"/>`),
		g.Group(children),
	)
}