package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarCrashSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20H3v-9l2.45-7h5.85q-.15.475-.225.975T11 6H6.85L5.8 9h5.875q.55 1.2 1.538 2.087t2.187 1.388q-.2.225-.3.488t-.1.537q0 .625.438 1.063T16.5 15q.625 0 1.063-.438T18 13.5q0-.125-.013-.25t-.062-.25q.8 0 1.575-.163t1.5-.512V20h-3v-2H6v2Zm1.5-5q.625 0 1.063-.438T9 13.5q0-.625-.438-1.063T7.5 12q-.625 0-1.063.438T6 13.5q0 .625.438 1.063T7.5 15ZM18 11q-2.075 0-3.538-1.463T13 6q0-2.05 1.45-3.525T18 1q2.075 0 3.538 1.462T23 6q0 2.075-1.463 3.538T18 11Zm-.5-4h1V3h-1v4Zm0 2h1V8h-1v1Z"/>`),
		g.Group(children),
	)
}