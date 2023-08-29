package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-2h2.75l-.4-.35q-1.3-1.15-1.825-2.625T3 12.05q0-2.775 1.662-4.938T9 4.25v2.1Q7.2 7 6.1 8.562T5 12.05q0 1.125.425 2.188T6.75 16.2l.25.25V14h2v6H3Zm17.925-9H18.9q-.125-.875-.537-1.675T17.25 7.8L17 7.55V10h-2V4h6v2h-2.75l.4.35q1.025 1.05 1.575 2.225t.7 2.425ZM17 23l-.3-1.5q-.3-.125-.563-.263T15.6 20.9l-1.45.45l-1-1.7l1.15-1q-.05-.35-.05-.65t.05-.65l-1.15-1l1-1.7l1.45.45q.275-.2.538-.337t.562-.263L17 13h2l.3 1.5q.3.125.563.275t.537.375l1.45-.5l1 1.75l-1.15 1q.05.3.05.625t-.05.625l1.15 1l-1 1.7l-1.45-.45q-.275.2-.537.338t-.563.262L19 23h-2Zm1-3q.825 0 1.413-.588T20 18q0-.825-.588-1.413T18 16q-.825 0-1.413.588T16 18q0 .825.588 1.413T18 20Z"/>`),
		g.Group(children),
	)
}