package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PayCodeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="10" height="10" x="32" y="6" fill="#2F88FF"/><rect width="10" height="10" x="32" y="32" fill="#2F88FF"/><rect width="10" height="10" x="6" y="32" fill="#2F88FF"/><rect width="10" height="10" x="6" y="6" fill="#2F88FF"/><path d="M8 24L30 24"/><path d="M38 24L40 24"/><path d="M24 37V39"/><path d="M24 17V31"/><path d="M24 8V10"/></g>`),
		g.Group(children),
	)
}