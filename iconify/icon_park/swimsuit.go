package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swimsuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M14 4V11"/><path d="M34 4V11"/><path d="M12 31H36V36C33 36 26 39 26 44H21C21 39 15 36 12 36V31Z"/><circle cx="14" cy="18" r="7" fill="#2F88FF"/><circle cx="34" cy="18" r="7" fill="#2F88FF"/><path d="M21 18H27"/></g>`),
		g.Group(children),
	)
}