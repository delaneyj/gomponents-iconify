package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="38" height="38" x="5" y="5" stroke="currentColor" stroke-width="4" rx="3"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M11 12h8v24h-8zm14 0h12v6H25z"/><circle cx="25" cy="24" r="2" fill="currentColor"/><circle cx="25" cy="30" r="2" fill="currentColor"/><circle cx="25" cy="36" r="2" fill="currentColor"/><circle cx="31" cy="24" r="2" fill="currentColor"/><circle cx="31" cy="30" r="2" fill="currentColor"/><circle cx="31" cy="36" r="2" fill="currentColor"/><circle cx="37" cy="24" r="2" fill="currentColor"/><circle cx="37" cy="30" r="2" fill="currentColor"/><circle cx="37" cy="36" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}