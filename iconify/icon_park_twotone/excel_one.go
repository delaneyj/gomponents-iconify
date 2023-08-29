package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExcelOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTExcelOne0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#555" rx="3"/><path stroke-linecap="round" d="M30 16H18m0 16V16m10 8H18m12 8H18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTExcelOne0)"/>`),
		g.Group(children),
	)
}