package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FTwoKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" d="M26 20C26 17.4667 28.1334 16 30 16C31.8667 16 34 17.4667 34 20C34 24.56 26 27.9466 26 32H34"/><path stroke="#fff" d="M21 16H14V32"/><path stroke="#fff" d="M14 24H21"/></g>`),
		g.Group(children),
	)
}