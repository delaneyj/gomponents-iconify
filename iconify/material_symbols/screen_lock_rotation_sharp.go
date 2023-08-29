package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenLockRotationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4h2V3q0-.425-.288-.713T18 2q-.425 0-.713.288T17 3v1Zm-4 18.95q-2.475 0-4.663-.95t-3.812-2.575Q2.9 17.8 1.95 15.612T1 10.95h2q0 1.8.613 3.425T5.3 17.3q1.075 1.3 2.55 2.212t3.175 1.263L8.4 18.15l1.4-1.4l5.9 5.9q-.675.15-1.325.225T13 22.95Zm1.75-3.2L4.2 9.2l7.05-7.05L13 3.9V11h7.075l1.775 1.65l-7.1 7.1ZM15 9V4h1V3q0-.825.588-1.413T18 1q.825 0 1.413.588T20 3v1h1v5h-6Z"/>`),
		g.Group(children),
	)
}