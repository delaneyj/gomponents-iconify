package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdrAutoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.8 14.2h4.4l.8 2.225q.1.275.325.425t.5.15q.475 0 .738-.388t.087-.812l-3.425-9.2q-.1-.275-.337-.437T12.35 6h-.7q-.3 0-.537.163t-.338.437L7.35 15.825q-.175.425.088.8t.737.375q.275 0 .5-.162T9 16.4l.8-2.2Zm.55-1.6l1.6-4.55h.1l1.6 4.55h-3.3ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}