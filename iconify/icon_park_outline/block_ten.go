package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockTen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 6h12v12H24V6Z"/><path d="M24 6h12v12H24V6ZM12 6h12v12H12V6Zm0 24h12v12H12V30Zm0-12h12v12H12V18Z"/></g>`),
		g.Group(children),
	)
}