package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullDressLonguette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTFullDressLonguette0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m18 18l-4-8h20l-4 8v6l10.5 20H7l11-20v-6Z"/><path d="M20 4v6m8-6v6M18 21h12m-12-2v4m12-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTFullDressLonguette0)"/>`),
		g.Group(children),
	)
}