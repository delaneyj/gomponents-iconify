package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalDiningRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.725 20.3q-.275-.275-.275-.7t.275-.7l9.55-9.55q-.45-1.05-.125-2.375T14.575 4.6q1.325-1.325 2.95-1.55t2.65.8q1.025 1.025.8 2.65t-1.55 2.95q-1.05 1.1-2.375 1.425t-2.375-.125L13.425 12l6.9 6.9q.275.275.275.7t-.275.7q-.3.3-.7.3t-.7-.3l-6.9-6.85l-6.9 6.85q-.3.3-.7.3t-.7-.3Zm3.65-7.85l-3-3Q3.25 8.325 3.062 6.825T3.625 4q.25-.425.75-.475t.875.35l5.325 5.375l-3.2 3.2Z"/>`),
		g.Group(children),
	)
}