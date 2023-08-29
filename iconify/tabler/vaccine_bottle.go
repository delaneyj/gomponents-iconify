package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zm1 2v.98c0 .877-.634 1.626-1.5 1.77c-.866.144-1.5.893-1.5 1.77V19a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-8.48c0-.877-.634-1.626-1.5-1.77A1.795 1.795 0 0 1 14 6.98V6m-7 6h10M7 18h10m-6-3h2"/>`),
		g.Group(children),
	)
}