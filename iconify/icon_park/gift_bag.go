package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="36" height="28" x="6" y="14" stroke-linejoin="round" rx="3"/><path fill="#2F88FF" stroke-linejoin="round" d="M6 32H42V39C42 40.6569 40.6569 42 39 42H9C7.34315 42 6 40.6569 6 39V32Z"/><circle cx="19" cy="9" r="5" fill="#2F88FF"/><circle cx="28" cy="10" r="4" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 20L24 14L31 20"/></g>`),
		g.Group(children),
	)
}