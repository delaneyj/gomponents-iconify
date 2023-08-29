package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdentityPlatformOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13q1.45 0 2.475-1.025Q15.5 10.95 15.5 9.5q0-1.45-1.025-2.475Q13.45 6 12 6q-1.45 0-2.475 1.025Q8.5 8.05 8.5 9.5q0 1.45 1.025 2.475Q10.55 13 12 13Zm0-2q-.625 0-1.062-.438q-.438-.437-.438-1.062t.438-1.062Q11.375 8 12 8t1.062.438q.438.437.438 1.062t-.438 1.062Q12.625 11 12 11Zm0 11.5L3 17V7l9-5.5L21 7v10Zm0-2.35l3.675-2.25q-.825-.425-1.75-.662Q13 17 12 17t-1.925.238q-.925.237-1.75.662Zm-5.6-3.425q1.2-.825 2.613-1.275Q10.425 15 12 15q1.575 0 2.988.45q1.412.45 2.612 1.275l1.4-.85v-7.75L12 3.85L5 8.125v7.75Zm5.6-6.45Z"/>`),
		g.Group(children),
	)
}