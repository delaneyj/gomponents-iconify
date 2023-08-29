package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsSystemDaydreamOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16h6.5q1.05 0 1.775-.725T18 13.5q0-1.05-.725-1.775T15.5 11h-.05q-.2-1.275-1.113-2.138T12.15 8q-1.05 0-1.95.525T8.85 10h-.125q-1.175 0-1.95.9T6 13q0 1.25.875 2.125T9 16Zm0-2q-.425 0-.713-.288T8 13q0-.425.288-.713T9 12h1.25v-.25q0-.725.513-1.238T12 10q.725 0 1.238.513t.512 1.237V13h1.75q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15H9Zm-5 6q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}