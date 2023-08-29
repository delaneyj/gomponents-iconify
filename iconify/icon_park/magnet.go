package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 10V26C6 35.9411 14.0589 44 24 44C33.9411 44 42 35.9411 42 26V10M14 10V26C14 31.5228 18.4772 36 24 36C29.5228 36 34 31.5228 34 26V10"/><path fill="#2F88FF" d="M14 4H6V10H14V4Z"/><path fill="#2F88FF" d="M34 4H42V10H34V4Z"/></g>`),
		g.Group(children),
	)
}