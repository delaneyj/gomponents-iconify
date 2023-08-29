package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SheetsRtlRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.3 21.3l-2.6-2.6q-.3-.3-.3-.7t.3-.7l2.6-2.6q.3-.3.7-.3t.7.3q.3.3.288.713t-.288.712L6.825 17H19q.425 0 .713.288T20 18q0 .425-.288.713T19 19H6.825l.875.875q.275.3.287.713T7.7 21.3q-.3.3-.7.3t-.7-.3ZM5.5 13q-.625 0-1.062-.438T4 11.5v-9q0-.625.438-1.063T5.5 1h13q.625 0 1.063.438T20 2.5v9q0 .625-.438 1.063T18.5 13h-13ZM6 6h5V3H6v3Zm7 0h5V3h-5v3Zm-2 5V8H6v3h5Zm2 0h5V8h-5v3Z"/>`),
		g.Group(children),
	)
}