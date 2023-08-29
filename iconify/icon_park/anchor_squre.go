package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnchorSqure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 38C42 28.0589 33.9411 18 24 18C14.0589 18 6 28.0589 6 38"/><path d="M20 14L10 14"/><path d="M38 14H28"/><circle cx="24" cy="14" r="4" fill="#2F88FF"/><rect width="8" height="8" x="20" y="10" fill="#2F88FF"/><rect width="6" height="6" x="38" y="11" fill="#2F88FF"/><rect width="6" height="6" x="4" y="11" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}