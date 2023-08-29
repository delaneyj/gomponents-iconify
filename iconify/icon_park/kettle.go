package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kettle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="7" x="6.778" y="37" fill="#2F88FF"/><path d="M38.7778 36L36.7778 10H4.77783L11.3438 17.5493C11.8727 18.1574 12.1347 18.9527 12.0707 19.7561L10.7778 36"/><path d="M26.7778 18H21.7778"/><path d="M26.7778 24H21.7778"/><path d="M26.7778 30H21.7778"/><path d="M36.7778 10H44.7778V26H38.7778"/><path d="M19.7778 9V4H28.7778V9"/></g>`),
		g.Group(children),
	)
}