package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 20q-.625 0-1.063-.438T2 18.5q0-.625.438-1.063T3.5 17q.625 0 1.063.438T5 18.5q0 .625-.438 1.063T3.5 20ZM8 20q-.4 0-.7-.238t-.375-.637Q6.65 17.55 5.513 16.45T2.8 15.075q-.375-.05-.587-.362T2 14q0-.425.263-.713t.637-.237q2.35.3 4.025 1.975t2 4.025q.05.4-.225.675T8 20Zm4 0q-.425 0-.713-.275t-.337-.7q-.35-3.2-2.612-5.425t-5.463-2.55q-.4-.05-.638-.35T2 10q0-.425.263-.725t.637-.25q4.025.325 6.85 3.138t3.2 6.812q.05.425-.237.725T12 20Zm3 0q0-2.725-1.012-5.088t-2.775-4.125Q9.45 9.025 7.075 8.014T2 7V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20h-5Z"/>`),
		g.Group(children),
	)
}