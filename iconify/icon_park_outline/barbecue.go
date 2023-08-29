package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barbecue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="8" height="8" x="12" y="4" stroke-linejoin="round" rx="4"/><rect width="8" height="8" x="12" y="22" stroke-linejoin="round" rx="4"/><path d="M16 31v13"/><rect width="8" height="8" x="28" y="4" stroke-linejoin="round" rx="4"/><rect width="8" height="8" x="28" y="22" stroke-linejoin="round" rx="4"/><path d="M32 31v13"/><path stroke-linejoin="round" stroke-miterlimit="2" d="M13 17h6m10 0h6M13 36h6m10 0h6"/><path d="M16 14v6m16-6v6"/></g>`),
		g.Group(children),
	)
}