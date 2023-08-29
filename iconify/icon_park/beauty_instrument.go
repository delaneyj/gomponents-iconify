package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeautyInstrument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#000" d="M17 34L31 34"/><path stroke="#000" d="M16 27L32 27"/><rect width="30" height="16" x="9" y="4" fill="#2F88FF" stroke="#000" rx="4"/><path stroke="#000" d="M13.9999 20L18.9999 44H28.9999L33.9999 20"/><path stroke="#fff" d="M17 12L30 12"/></g>`),
		g.Group(children),
	)
}