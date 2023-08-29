package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refrigerator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRefrigerator0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="28" height="36" x="9" y="4" fill="#fff" stroke="#fff" rx="2"/><path stroke="#000" stroke-linecap="round" d="M9 22h28"/><path stroke="#fff" stroke-linecap="round" d="M9 20v4m28-4v4"/><path stroke="#000" stroke-linecap="round" d="M15 29v4m0-22v4"/><path stroke="#fff" stroke-linecap="round" d="M33 40v4m-20-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRefrigerator0)"/>`),
		g.Group(children),
	)
}