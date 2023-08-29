package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 38C42 28.0589 33.9411 18 24 18C14.0589 18 6 28.0589 6 38"/><path d="M20 14L10 14"/><path d="M38 14H28"/><circle cx="7" cy="14" r="3" fill="#2F88FF"/><circle cx="41" cy="14" r="3" fill="#2F88FF"/><circle cx="24" cy="14" r="4" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}