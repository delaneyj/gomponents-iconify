package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" stroke-linecap="round" d="M6 28L24 28"/><path stroke="#fff" stroke-linecap="round" d="M24 20H42"/><path stroke="#000" stroke-linecap="round" d="M6 25V31"/><path stroke="#000" stroke-linecap="round" d="M42 17V23"/><path stroke="#fff" stroke-linecap="round" d="M24 42V6"/><path stroke="#000" stroke-linecap="round" d="M21 6H27"/><path stroke="#000" stroke-linecap="round" d="M21 42H27"/></g>`),
		g.Group(children),
	)
}