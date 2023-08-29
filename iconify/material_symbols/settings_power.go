package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsPower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 12V2h2v10h-2Zm1 7q-3.35 0-5.675-2.325T4 11q0-1.975.912-3.7T7.45 4.45L8.9 5.9q-1.35.8-2.125 2.163T6 11q0 2.5 1.75 4.25T12 17q2.5 0 4.25-1.75T18 11q0-1.575-.775-2.938T15.1 5.9l1.45-1.45q1.625 1.125 2.538 2.85T20 11q0 3.35-2.325 5.675T12 19Zm-4 5q-.425 0-.713-.288T7 23q0-.425.288-.713T8 22q.425 0 .713.288T9 23q0 .425-.288.713T8 24Zm4 0q-.425 0-.713-.288T11 23q0-.425.288-.713T12 22q.425 0 .713.288T13 23q0 .425-.288.713T12 24Zm4 0q-.425 0-.713-.288T15 23q0-.425.288-.713T16 22q.425 0 .713.288T17 23q0 .425-.288.713T16 24Z"/>`),
		g.Group(children),
	)
}