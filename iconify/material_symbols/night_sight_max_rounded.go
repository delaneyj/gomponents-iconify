package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightSightMaxRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 8h-2q-.425 0-.713-.288T14 7q0-.425.288-.713T15 6h2V4q0-.425.288-.713T18 3q.425 0 .713.288T19 4v2h2q.425 0 .713.288T22 7q0 .425-.288.713T21 8h-2v2q0 .425-.288.713T18 11q-.425 0-.713-.288T17 10V8Zm1.35 8.175q-.95 2.2-2.938 3.513T11 21q-3.35 0-5.675-2.325T3 13q0-2.875 1.837-5.075T9.5 5.15q.3-.05.537.05t.363.325q.125.225.15.537t-.1.663q-.225.55-.337 1.125T10 9q0 2.5 1.75 4.25T16 15q.275 0 .525-.013t.5-.062q.45-.075.738.025t.462.3q.175.2.212.438t-.087.487Z"/>`),
		g.Group(children),
	)
}