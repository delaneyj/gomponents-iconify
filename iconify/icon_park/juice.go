package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Juice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M15 24H33L31.2 44H16.8L15 24Z"/><rect width="26" height="6" x="11" y="18" rx="3"/><path fill="#2F88FF" d="M24 8C18.4772 8 14 12.4772 14 18H34C34 12.4772 29.5228 8 24 8Z"/><path stroke-linecap="round" d="M28 4L26 8"/></g>`),
		g.Group(children),
	)
}