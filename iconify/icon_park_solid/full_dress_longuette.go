package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullDressLonguette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFullDressLonguette0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="m18 18l-4-8h20l-4 8v6l10.5 20H7l11-20v-6Z"/><path stroke="#fff" d="M20 4v6m8-6v6"/><path stroke="#000" d="M18 21h12"/><path stroke="#fff" d="M18 19v4m12-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFullDressLonguette0)"/>`),
		g.Group(children),
	)
}