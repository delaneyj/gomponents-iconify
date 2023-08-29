package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OverallReduction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOverallReduction0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M15 15h18v18H15V15Z"/><path d="M11 43v-6H5m32 6v-6h6M11 5v6H5m32-6v6h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOverallReduction0)"/>`),
		g.Group(children),
	)
}