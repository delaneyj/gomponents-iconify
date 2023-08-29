package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4.5 18L24 8L44 18"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 18V4"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 30V16"/><circle cx="10" cy="36" r="6" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 30V16"/><path fill="#2F88FF" stroke-linecap="round" stroke-linejoin="round" d="M32 36L38 30L44 36L38 42L32 36Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 23V18H19V23"/></g>`),
		g.Group(children),
	)
}