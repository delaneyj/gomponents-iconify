package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barbecue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><rect width="8" height="8" x="12" y="4" fill="#2F88FF" stroke-linejoin="round" rx="4"/><rect width="8" height="8" x="12" y="22" fill="#2F88FF" stroke-linejoin="round" rx="4"/><line x1="16" x2="16" y1="31" y2="44"/><rect width="8" height="8" x="28" y="4" fill="#2F88FF" stroke-linejoin="round" rx="4"/><rect width="8" height="8" x="28" y="22" fill="#2F88FF" stroke-linejoin="round" rx="4"/><line x1="32" x2="32" y1="31" y2="44"/><path stroke-linejoin="round" stroke-miterlimit="2" d="M13 17H19"/><path stroke-linejoin="round" stroke-miterlimit="2" d="M29 17H35"/><path stroke-linejoin="round" stroke-miterlimit="2" d="M13 36H19"/><path stroke-linejoin="round" stroke-miterlimit="2" d="M29 36H35"/><line x1="16" x2="16" y1="14" y2="20"/><line x1="32" x2="32" y1="14" y2="20"/></g>`),
		g.Group(children),
	)
}