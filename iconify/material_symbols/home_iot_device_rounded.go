package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeIotDeviceRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.025 10h1V6.5q0-.225-.138-.363T5.526 6q-.225 0-.363.138t-.137.362V10Zm.5 10q-1.875 0-3.188-1.3t-1.312-3.2q0-1.2.55-2.15t1.45-1.55V6.5q0-1.05.725-1.775T5.525 4q1.05 0 1.775.725T8.025 6.5v5.3q.9.6 1.45 1.55t.55 2.15q0 1.875-1.312 3.188T5.525 20Zm8.175-5q-1.2-.825-1.95-2.125T11 10q0-2.5 1.75-4.25T17 4q2.5 0 4.25 1.75T23 10q0 1.575-.75 2.875T20.3 15h-6.6Zm3.3 5q-.6 0-1.05-.45t-.45-1.05h3q0 .6-.45 1.05T17 20Zm-2-2q-.425 0-.713-.288T14 17q0-.425.288-.713T15 16h4q.425 0 .713.288T20 17q0 .425-.288.713T19 18h-4Z"/>`),
		g.Group(children),
	)
}