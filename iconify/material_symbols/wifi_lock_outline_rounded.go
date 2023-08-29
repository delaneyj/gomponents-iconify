package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiLockOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21q-.425 0-.713-.288T16 20v-3q0-.425.288-.713T17 16v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1q.425 0 .713.288T22 17v3q0 .425-.288.713T21 21h-4Zm1-5h2v-1q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1Zm-6.7 4.3L.75 9.75q-.3-.3-.3-.738T.775 8.3q2.35-2.1 5.2-3.2T12 4q3.125 0 6.025 1.088t5.2 3.212q.325.275.325.713t-.3.737l-1.775 1.775q-.525-.25-1.113-.375T19.15 11l1.95-1.95q-1.975-1.5-4.3-2.275T12 6q-2.475 0-4.8.775T2.9 9.05l9.1 9.1l1-1q.025.625.15 1.213t.375 1.112l-.825.825q-.3.3-.7.3t-.7-.3Zm.7-8.225Z"/>`),
		g.Group(children),
	)
}