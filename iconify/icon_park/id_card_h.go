package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCardH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="40" height="32" x="4" y="8" rx="2"/><path fill="#2F88FF" d="M17 25C19.2091 25 21 23.2091 21 21C21 18.7909 19.2091 17 17 17C14.7909 17 13 18.7909 13 21C13 23.2091 14.7909 25 17 25Z"/><path stroke-linecap="round" d="M23 31C23 27.6863 20.3137 25 17 25C13.6863 25 11 27.6863 11 31"/><path stroke-linecap="round" d="M28 20H36"/><path stroke-linecap="round" d="M30 28H36"/></g>`),
		g.Group(children),
	)
}