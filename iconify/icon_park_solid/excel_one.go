package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExcelOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSExcelOne0"><g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path stroke="#000" stroke-linecap="round" d="M30 16H18m0 16V16m10 8H18m12 8H18"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSExcelOne0)"/>`),
		g.Group(children),
	)
}