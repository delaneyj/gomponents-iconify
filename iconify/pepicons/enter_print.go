package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><rect width="10" height="15" x="8.25" y="4" fill="currentColor" opacity=".8" rx="1"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.867 7.5L10.95 10L8.867 7.5Zm0 5L10.95 10l-2.083 2.5Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" d="M10.5 10H3m12.5 7H6.3m9.2-14H6.326M6.25 17v-4m0-6V3m9.35 14V3"/></g>`),
		g.Group(children),
	)
}