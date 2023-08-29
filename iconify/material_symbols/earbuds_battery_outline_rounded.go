package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarbudsBatteryOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 18q-.425 0-.713-.288T16 17V8q0-.425.288-.713T17 7h1v-.5q0-.2.15-.35T18.5 6h1q.2 0 .35.15t.15.35V7h1q.425 0 .713.288T22 8v9q0 .425-.288.713T21 18h-4Zm1-2h2h-2ZM5.375 18q-1.425 0-2.4-.975T2 14.625V8q0-.825.588-1.413T4 6q.825 0 1.413.588T6 8q0 .825-.588 1.413T4 10q-.125 0-.25-.025T3.5 9.9v4.725q0 .8.537 1.338t1.338.537q.8 0 1.338-.537t.537-1.338v-5.25q0-1.425.975-2.4t2.4-.975q1.425 0 2.4.975t.975 2.4V16q0 .825-.588 1.413T12 18q-.825 0-1.413-.588T10 16q0-.825.588-1.413T12 14q.125 0 .25.025t.25.075V9.375q0-.8-.537-1.338T10.625 7.5q-.8 0-1.338.537T8.75 9.375v5.25q0 1.425-.975 2.4t-2.4.975ZM18 16h2V9h-2v7Z"/>`),
		g.Group(children),
	)
}