package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageSquareDashed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6V5c0-1.1.9-2 2-2h2m4 0h3m4 0h1c1.1 0 2 .9 2 2m0 4v2m0 4c0 1.1-.9 2-2 2h-1m-4 0h-3m-4 0l-4 4v-5m0-4v-2"/>`),
		g.Group(children),
	)
}