package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 16L12.5 17.5"/><path d="M8 42L14 24"/><path d="M38 42L21 24"/><path d="M42 16L22.5 17.5"/><path d="M20 6L18 13"/></g>`),
		g.Group(children),
	)
}