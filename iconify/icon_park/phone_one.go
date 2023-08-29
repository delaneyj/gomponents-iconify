package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="38" height="38" x="5" y="5" stroke="#000" stroke-width="4" rx="3"/><rect width="8" height="24" x="11" y="12" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><rect width="12" height="6" x="25" y="12" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><circle cx="25" cy="24" r="2" fill="#000"/><circle cx="25" cy="30" r="2" fill="#000"/><circle cx="25" cy="36" r="2" fill="#000"/><circle cx="31" cy="24" r="2" fill="#000"/><circle cx="31" cy="30" r="2" fill="#000"/><circle cx="31" cy="36" r="2" fill="#000"/><circle cx="37" cy="24" r="2" fill="#000"/><circle cx="37" cy="30" r="2" fill="#000"/><circle cx="37" cy="36" r="2" fill="#000"/></g>`),
		g.Group(children),
	)
}