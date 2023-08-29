package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screwdriver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSScrewdriver0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M19 24h10v5c.961.976 2.039 2.524 3 3.5V44H16V32.5l3-3.5v-5Z"/><path stroke-linecap="round" d="M21 13v11h6V13l2-3l-2-6h-6l-2 6l2 3Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSScrewdriver0)"/>`),
		g.Group(children),
	)
}