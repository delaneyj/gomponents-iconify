package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 7c4.418 0 8 .895 8 2s-3.582 2-8 2s-8-.895-8-2c0-.357.375-.693 1.033-.984"/><path d="M3 9v9.343c0 1.061.44 2.08 1.409 2.513C5.624 21.399 7.711 22 11 22c3.29 0 5.377-.601 6.591-1.144c.968-.433 1.409-1.452 1.409-2.513V9m0 1v0a3 3 0 0 1 3 3v0a3 3 0 0 1-3 3v0M7 3v4m4-5v2m4 0v3"/></g>`),
		g.Group(children),
	)
}