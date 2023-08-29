package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.124 20.119a3 3 0 1 0-4.248-4.237a3 3 0 0 0 4.248 4.237Zm0 0L22 22M7 2H4v3m0 6v2m7-11h2m-2 20h2m7-11v2M17 2h3v3M7 22H4v-3"/>`),
		g.Group(children),
	)
}