package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchAppSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 22H9.025l-6.35-8.075l1.6-1.675L7.5 14.2V5h2.025v6H20v11ZM4.175 8.5q-.325-.55-.5-1.188T3.5 6q0-2.075 1.463-3.538T8.5 1q2.075 0 3.538 1.463T13.5 6q0 .675-.175 1.313t-.5 1.187l-1.725-1q.2-.35.3-.713T11.5 6q0-1.25-.875-2.125T8.5 3q-1.25 0-2.125.875T5.5 6q0 .425.1.788t.3.712l-1.725 1Z"/>`),
		g.Group(children),
	)
}