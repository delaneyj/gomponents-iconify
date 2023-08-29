package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StormOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.15 22q.625-1.55.875-3.175t.125-3.275q-.975 2.075-2.912 3.262T12 20q-2.1 0-3.8-.988t-2.9-2.637q-1.2-1.65-1.85-3.788T2.8 8.15q0-1.575.213-3.113T3.75 2h2.1q-.6 1.55-.863 3.175T4.85 8.45q.975-2.075 2.913-3.262T12 4q2.1 0 3.8.988t2.9 2.637q1.2 1.65 1.85 3.788t.65 4.437q0 1.575-.212 3.113T20.25 22h-2.1ZM12 18q2.5 0 4.25-1.75T18 12q0-2.5-1.75-4.25T12 6Q9.5 6 7.75 7.75T6 12q0 2.5 1.75 4.25T12 18Zm0-2q1.65 0 2.825-1.175T16 12q0-1.65-1.175-2.825T12 8q-1.65 0-2.825 1.175T8 12q0 1.65 1.175 2.825T12 16Zm0-2q-.825 0-1.413-.588T10 12q0-.825.588-1.413T12 10q.825 0 1.413.588T14 12q0 .825-.588 1.413T12 14Zm0-2Z"/>`),
		g.Group(children),
	)
}