package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TitleLevel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 8.00049V40.0005"/><path stroke-linejoin="round" d="M24 8.00049V40.0005"/><path stroke-linejoin="round" d="M7 24.0005H23"/><path d="M32 24V40"/><path d="M32 31.0239C32 28.4599 34 26.0005 37 26.0005C40 26.0005 42 28.3585 42 31.0239C42 32.8009 42 36.4642 42 40.0139"/></g>`),
		g.Group(children),
	)
}