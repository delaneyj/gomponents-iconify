package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenLockRotationOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22.95q-2.475 0-4.663-.95t-3.812-2.575Q2.9 17.8 1.95 15.612T1 10.95h2q0 1.8.613 3.425T5.3 17.3q1.075 1.3 2.55 2.212t3.175 1.263L8.4 18.15l1.4-1.4l5.9 5.9q-.675.15-1.325.225T13 22.95ZM15 9V4h1V3q0-.825.588-1.413T18 1q.825 0 1.413.588T20 3v1h1v5h-6Zm2-5h2V3q0-.425-.288-.712T18 2q-.425 0-.713.288T17 3v1Zm.95 7.65l1.4-1.4l2.45 2.45l-7.05 7.05L4.2 9.2l7.05-7.05L13.7 4.6l-1.35 1.45l-1.1-1.1L7 9.2l7.75 7.75L19 12.7l-1.05-1.05Zm-4.95-.7Z"/>`),
		g.Group(children),
	)
}