package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="38" height="14" x="5" y="28" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><rect width="38" height="14" x="5" y="6" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><rect width="4" height="4" x="11" y="11" fill="#fff" rx="2"/><rect width="4" height="4" x="11" y="33" fill="#fff" rx="2"/><rect width="4" height="4" x="19" y="11" fill="#fff" rx="2"/><rect width="4" height="4" x="19" y="33" fill="#fff" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 13H35"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 35H35"/></g>`),
		g.Group(children),
	)
}