package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scanning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3H3v3m-1 6h20M9 19v-4m3 1v-1m3 2v-2m-3 6v-3m6-15h3v3M6 21H3v-3m15 3h3v-3"/>`),
		g.Group(children),
	)
}