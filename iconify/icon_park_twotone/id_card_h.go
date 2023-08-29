package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCardH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTIdCardH0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" rx="2"/><path fill="#555" d="M17 25a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M23 31a6 6 0 0 0-12 0m17-11h8m-6 8h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTIdCardH0)"/>`),
		g.Group(children),
	)
}