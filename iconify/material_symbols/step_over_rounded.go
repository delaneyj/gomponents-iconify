package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepOverRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19q-1.25 0-2.125-.875T9 16q0-1.25.875-2.125T12 13q1.25 0 2.125.875T15 16q0 1.25-.875 2.125T12 19Zm5.2-9q-.8-1.35-2.163-2.175T12 7q-1.925 0-3.45 1.1t-2.175 2.825q-.15.425-.463.75T5.175 12q-.45 0-.725-.363t-.15-.787Q5 8.3 7.113 6.65T11.975 5q1.825 0 3.375.738T18 7.75V6q0-.425.288-.713T19 5q.425 0 .713.288T20 6v5q0 .425-.288.713T19 12h-5q-.425 0-.713-.288T13 11q0-.425.288-.713T14 10h3.2Z"/>`),
		g.Group(children),
	)
}