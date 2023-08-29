package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><circle cx="22.834" cy="22.834" r="17" fill="#2F88FF" stroke="#000" stroke-linejoin="round"/><path stroke="#000" d="M35 35L41 41"/><path stroke="#fff" stroke-linejoin="round" d="M23 17V31"/><path stroke="#fff" stroke-linejoin="round" d="M18 17H23H28"/></g>`),
		g.Group(children),
	)
}