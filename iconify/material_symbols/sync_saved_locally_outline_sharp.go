package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncSavedLocallyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.925 14.05L16.6 8.4l-1.425-1.425l-4.25 4.25L8.8 9.1l-1.4 1.4l3.525 3.55ZM1 21v-2h22v2H1Zm1-3V3h20v15H2Zm2-2h16V5H4v11Zm0 0V5v11Z"/>`),
		g.Group(children),
	)
}