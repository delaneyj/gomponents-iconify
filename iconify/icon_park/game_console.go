package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GameConsole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="28" height="40" x="10" y="4" stroke="#000" stroke-width="4" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 34H24"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 30V38"/><rect width="16" height="9" x="16" y="10" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"/><circle cx="31" cy="30" r="2" fill="#000"/><circle cx="31" cy="38" r="2" fill="#000"/></g>`),
		g.Group(children),
	)
}