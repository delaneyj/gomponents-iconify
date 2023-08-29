package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullDressLonguette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m18 18l-4-8h20l-4 8v6l10.5 20H7l11-20v-6Zm2-14v6m8-6v6M18 21h12m-12-2v4m12-4v4"/>`),
		g.Group(children),
	)
}