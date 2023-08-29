package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phonograph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="36" height="36" x="6" y="6" stroke="#000" stroke-width="4" rx="3"/><circle cx="24" cy="25" r="11" fill="#2F88FF" stroke="#000" stroke-width="4"/><rect width="4" height="4" x="22" y="23" fill="#fff" rx="2"/><rect width="4" height="4" x="34" y="34" fill="#000" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 20L36 12"/></g>`),
		g.Group(children),
	)
}