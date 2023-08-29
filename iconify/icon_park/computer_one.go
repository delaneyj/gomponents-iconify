package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComputerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-width="4" d="M10 6C10 4.89543 10.8954 4 12 4H36C37.1046 4 38 4.89543 38 6V42C38 43.1046 37.1046 44 36 44H12C10.8954 44 10 43.1046 10 42L10 6Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 12L31 12"/><circle cx="17" cy="21" r="2" fill="#000"/><circle cx="17" cy="27" r="2" fill="#000"/><circle cx="17" cy="33" r="2" fill="#000"/><circle cx="24" cy="21" r="2" fill="#000"/><circle cx="24" cy="27" r="2" fill="#000"/><circle cx="24" cy="33" r="2" fill="#000"/><circle cx="31" cy="21" r="2" fill="#000"/><circle cx="31" cy="27" r="2" fill="#000"/><circle cx="31" cy="33" r="2" fill="#000"/></g>`),
		g.Group(children),
	)
}