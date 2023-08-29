package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UfoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2.731 10c.876-.75 2.274-1.373 3.996-1.79m-3.996 5.442C4.355 15.042 7.774 16 11.73 16c5.524 0 10.002-1.869 10.002-4.174c0-1.167-1.148-2.223-3-2.98a13.96 13.96 0 0 0-2.001-.635M11.729 10c-3.191 0-4.388-.532-4.802-.82c-.146-.1-.2-.274-.2-.451A4.73 4.73 0 0 1 11.457 4H12a4.73 4.73 0 0 1 4.73 4.729c0 .177-.054.35-.2.452c-.281.195-.922.502-2.3.68M12 16v3m-6.5-3.5l-1 2m14-2l1 2"/><circle cx="12" cy="13" r="1" fill="currentColor"/><circle cx="7" cy="12" r="1" fill="currentColor"/><circle cx="17" cy="12" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}