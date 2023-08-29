package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpecialCharacterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 5Q9.7 5 8.1 6.6t-1.6 3.9q0 1.575.825 2.9t2.25 2.025q.6.3 1.012.813T11 17.4v2.1q0 .625-.438 1.063T9.5 21h-5q-.625 0-1.063-.438T3 19.5q0-.625.438-1.063T4.5 18H8q-2.1-1.125-3.3-3.125T3.5 10.5q0-3.55 2.475-6.025T12 2q3.55 0 6.025 2.475T20.5 10.5q0 2.375-1.2 4.375T16 18h3.5q.625 0 1.063.438T21 19.5q0 .625-.438 1.063T19.5 21h-5q-.625 0-1.063-.438T13 19.5v-2.1q0-.65.413-1.163t1.012-.812q1.425-.7 2.25-2.025t.825-2.9q0-2.3-1.6-3.9T12 5Z"/>`),
		g.Group(children),
	)
}