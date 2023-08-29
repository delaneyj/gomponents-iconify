package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5.459 2L1 6.015L2.338 7.5l4.46-4.015L5.457 2ZM11 8h2v4h3v2h-5V8Z"/><path fill-rule="evenodd" d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0Zm2 0a7 7 0 1 1 14 0a7 7 0 0 1-14 0Z" clip-rule="evenodd"/><path d="M18.541 2L23 6.015L21.662 7.5l-4.46-4.015L18.543 2Z"/></g>`),
		g.Group(children),
	)
}