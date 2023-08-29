package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passkey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 12q-1.65 0-2.825-1.175T7 8q0-1.65 1.175-2.825T11 4q1.65 0 2.825 1.175T15 8q0 1.65-1.175 2.825T11 12Zm7.25 9L17 19.5v-3.675q-.9-.35-1.45-1.112T15 13q0-1.25.875-2.125T18 10q1.25 0 2.125.875T21 13q0 .95-.55 1.713T19 15.825V16l1 1l-1 1l1 1l-1.75 2ZM18 14.5q.625 0 1.063-.438T19.5 13q0-.625-.438-1.063T18 11.5q-.625 0-1.063.438T16.5 13q0 .625.438 1.063T18 14.5Zm-4.975-1.35q.05 1.15.537 2.163T15 16.974V20H3v-2.775q0-.85.425-1.575t1.175-1.1q1.5-.75 3.113-1.15T11 13q.5 0 1.012.038t1.013.112Z"/>`),
		g.Group(children),
	)
}