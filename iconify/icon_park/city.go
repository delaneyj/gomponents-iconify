package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func City(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 42H44"/><rect width="8" height="16" x="8" y="26" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" rx="2"/><path stroke="#fff" stroke-linecap="square" stroke-linejoin="round" stroke-width="4" d="M12 34H13"/><rect width="24" height="38" x="16" y="4" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" rx="2"/><rect width="4" height="4" x="22" y="10" fill="#fff"/><rect width="4" height="4" x="30" y="10" fill="#fff"/><rect width="4" height="4" x="22" y="17" fill="#fff"/><rect width="4" height="4" x="30" y="17" fill="#fff"/><rect width="4" height="4" x="30" y="24" fill="#fff"/><rect width="4" height="4" x="30" y="31" fill="#fff"/></g>`),
		g.Group(children),
	)
}