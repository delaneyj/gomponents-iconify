package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" d="M34 21L37 24L34 27"/><path stroke="#fff" d="M14 21L11 24L14 27"/><path stroke="#fff" d="M27 14L24 11L21 14"/><path stroke="#fff" d="M27 34L24 37L21 34"/></g>`),
		g.Group(children),
	)
}