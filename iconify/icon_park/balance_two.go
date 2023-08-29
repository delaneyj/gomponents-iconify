package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BalanceTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M16 22L10 12L4 22"/><path fill="#2F88FF" fill-rule="evenodd" d="M10 28C13.3137 28 16 25.3137 16 22H4C4 25.3137 6.68629 28 10 28Z" clip-rule="evenodd"/><path d="M44 22L38 12L32 22"/><path fill="#2F88FF" fill-rule="evenodd" d="M38 28C41.3137 28 44 25.3137 44 22H32C32 25.3137 34.6863 28 38 28Z" clip-rule="evenodd"/><path d="M24 6V42"/><path d="M10 12H24H38"/><path d="M10 12H24H38"/><path d="M38 42H24H10"/></g>`),
		g.Group(children),
	)
}