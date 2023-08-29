package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyLiraRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21q-.425 0-.713-.288T9 20v-4.15l-1.475.925q-.5.325-1.012.025T6 15.9q0-.275.125-.5t.35-.35L9 13.475v-2.35l-1.475.925q-.5.3-1.012.025T6 11.2q0-.275.125-.5t.325-.35L9 8.75V4q0-.425.288-.713T10 3q.425 0 .713.288T11 4v3.5l2.475-1.55q.5-.325 1.012-.038T15 6.8q0 .275-.125.5t-.35.35L11 9.875v2.35l2.475-1.55q.5-.3 1.012-.025t.513.875q0 .275-.125.5t-.35.35L11 14.6V19q1.85 0 3.225-1.175t1.7-2.925q.075-.4.35-.65t.65-.25q.45 0 .75.325t.25.75q-.375 2.55-2.325 4.237T11 21h-1Z"/>`),
		g.Group(children),
	)
}