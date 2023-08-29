package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capricornus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M18 11C18 7.68629 15.3137 5 12 5C8.68629 5 6 7.68629 6 11"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 11V29"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 11C30 7.68629 27.3137 5 24 5C20.6863 5 18 7.68629 18 11"/><path stroke-linecap="round" stroke-linejoin="round" d="M30 11V31V35.75C30 35.75 30 43 22 43"/><circle cx="36" cy="30" r="6" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}