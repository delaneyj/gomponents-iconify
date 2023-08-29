package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCardV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSIdCardV0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="32" height="40" x="8" y="4" rx="2"/><path fill="#fff" d="M24 19a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M30 25a6 6 0 0 0-12 0m0 6h12m-12 6h7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSIdCardV0)"/>`),
		g.Group(children),
	)
}