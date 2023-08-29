package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blockchain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 30V15L27.5 7.96875M20.5 7.96875L8 15V30M11 34.6875L24 42L32 37.5L37 34.6875"/><path fill="#2F88FF" d="M21 18.75L18 20.5V24V27.5L21 29.25L24 31L27 29.25L30 27.5V24V20.5L27 18.75L24 17L21 18.75Z"/><path d="M24 17V10"/><path d="M30 27L37 31"/><path d="M18 27L11 31"/><circle cx="24" cy="7" r="3" fill="#2F88FF"/><circle cx="8" cy="33" r="3" fill="#2F88FF"/><circle cx="40" cy="33" r="3" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}