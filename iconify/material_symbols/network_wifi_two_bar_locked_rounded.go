package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkWifiTwoBarLockedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 17.025q0 .65.125 1.263t.4 1.187l-.825.825q-.3.3-.7.3t-.7-.3L.75 9.75q-.3-.3-.3-.738T.8 8.3Q3.225 6.2 6 5.1T12 4q3.225 0 6 1.088T23.2 8.3q.325.275.338.713t-.288.737l-1.775 1.775q-.55-.25-1.137-.375T19.15 11l1.95-1.95q-1.975-1.5-4.3-2.275T12 6q-2.475 0-4.8.775T2.9 9.05l4.3 4.3q1.1-.65 2.312-1T12 12q.775 0 1.538.163t1.487.362q-.95.875-1.488 2.038T13 17.024ZM17 21q-.425 0-.713-.287T16 20v-3q0-.425.288-.713T17 16v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1q.425 0 .713.288T22 17v3q0 .425-.288.713T21 21h-4Zm1-5h2v-1q0-.425-.288-.712T19 14q-.425 0-.713.288T18 15v1Z"/>`),
		g.Group(children),
	)
}