package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3H3v3m15-3h3v3M6 21H3v-3m4 0v-1a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v1m-5-6a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm6 9h3v-3"/>`),
		g.Group(children),
	)
}