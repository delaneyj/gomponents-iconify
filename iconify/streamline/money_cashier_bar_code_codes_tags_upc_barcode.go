package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyCashierBarCodeCodesTagsUpcBarcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5.5v10m3-10v10m3-10v10M11 .5v10m2.5-10v10m-13 3h13"/>`),
		g.Group(children),
	)
}