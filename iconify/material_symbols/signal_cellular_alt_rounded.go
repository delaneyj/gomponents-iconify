package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.5 20q-.625 0-1.063-.438T17 18.5v-13q0-.625.438-1.063T18.5 4q.625 0 1.063.438T20 5.5v13q0 .625-.438 1.063T18.5 20Zm-12 0q-.3 0-.575-.113t-.488-.325q-.212-.212-.325-.487T5 18.5v-3q0-.625.438-1.063T6.5 14q.625 0 1.063.438T8 15.5v3q0 .3-.113.575t-.324.488q-.213.212-.488.325T6.5 20Zm6 0q-.625 0-1.063-.438T11 18.5v-8q0-.625.438-1.063T12.5 9q.625 0 1.063.438T14 10.5v8q0 .625-.438 1.063T12.5 20Z"/>`),
		g.Group(children),
	)
}