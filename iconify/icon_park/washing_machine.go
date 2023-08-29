package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashingMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="40" x="8.778" y="4" stroke="#000" stroke-width="4" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8.77783 15.5H40.7778"/><circle cx="28.778" cy="10" r="2" fill="#000"/><circle cx="34.778" cy="10" r="2" fill="#000"/><circle cx="24.778" cy="30" r="7" fill="#2F88FF" stroke="#000" stroke-width="4"/></g>`),
		g.Group(children),
	)
}