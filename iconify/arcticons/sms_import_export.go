package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsImportExport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.364 8.912H8.637A4.137 4.137 0 0 0 4.5 13.05v18.052a4.137 4.137 0 0 0 4.137 4.136h27.636l5.563 5.563a.975.975 0 0 0 1.664-.69V13.05a4.136 4.136 0 0 0-4.136-4.136Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 28.63h6.644a2.228 2.228 0 0 0 2.228-2.228V15.237"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m36.999 19.364l-4.127-4.127l-4.126 4.127M24 15.237h-6.644a2.229 2.229 0 0 0-2.228 2.229V28.63"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.001 24.504l4.127 4.127l4.126-4.127"/>`),
		g.Group(children),
	)
}