package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oscillator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 9V24C7 27.3137 9.68629 30 13 30H35C38.3137 30 41 27.3137 41 24V9"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M7 10C7 6.68629 9.68629 4 13 4H35C38.3137 4 41 6.68629 41 10C41 13.3137 38.3137 16 35 16H13C9.68629 16 7 13.3137 7 10Z"/><circle cx="15" cy="10" r="2" fill="#fff"/><circle cx="21" cy="10" r="2" fill="#fff"/><circle cx="27" cy="10" r="2" fill="#fff"/><circle cx="33" cy="10" r="2" fill="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M19 30V44"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29 30V44"/></g>`),
		g.Group(children),
	)
}