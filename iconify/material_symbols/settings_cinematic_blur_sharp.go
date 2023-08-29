package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsCinematicBlurSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 22v-7.85h.825l.675 1.35H16l-.75-1.5h1.5l.75 1.5H19l-.75-1.5h1.5l.75 1.5H22l-.75-1.5H23v8H13Zm1.5-1.5h7V17h-7v3.5ZM9.25 22l-.4-3.2q-.325-.125-.613-.3t-.562-.375L4.7 19.375l-2.75-4.75l2.575-1.95Q4.5 12.5 4.5 12.337v-.674q0-.163.025-.338L1.95 9.375l2.75-4.75l2.975 1.25q.275-.2.575-.375t.6-.3l.4-3.2h5.5l.4 3.2q.325.125.613.3t.562.375l2.975-1.25l2.75 4.75l-2.575 1.95q.025.175.025.338V12h-3.95q0-1.45-1.025-2.475T12.05 8.5q-1.45 0-2.475 1.025T8.55 12q0 1.2.675 2.1T11 15.35V22H9.25Z"/>`),
		g.Group(children),
	)
}