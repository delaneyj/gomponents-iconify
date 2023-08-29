package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Devices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 15H4V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2M2 18h12"/><path d="M14 9.2c0-.663.597-1.2 1.333-1.2h5.334C21.403 8 22 8.537 22 9.2v9.6c0 .663-.597 1.2-1.333 1.2h-5.334C14.597 20 14 19.463 14 18.8V9.2Zm4 7.8h.01"/></g>`),
		g.Group(children),
	)
}