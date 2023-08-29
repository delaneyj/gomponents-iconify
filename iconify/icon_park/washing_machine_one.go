package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashingMachineOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="40" x="8" y="4" stroke="#000" stroke-width="4" rx="2"/><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M8 12C8 13.1046 8.89543 14 10 14H38C39.1046 14 40 13.1046 40 12V6C40 4.89543 39.1046 4 38 4H10C8.89543 4 8 4.89543 8 6V12Z"/><circle cx="14" cy="9" r="2" fill="#fff"/><circle cx="20" cy="9" r="2" fill="#fff"/><circle cx="24" cy="29" r="7" fill="#2F88FF" stroke="#000" stroke-width="4"/></g>`),
		g.Group(children),
	)
}