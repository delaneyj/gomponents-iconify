package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorbellOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17.5q.425 0 .713-.288T13 16.5h-2q0 .425.288.713T12 17.5ZM8 16h8v-1h-1v-2.35q0-1.1-.6-2T12.75 9.5v-.25q0-.325-.213-.537T12 8.5q-.325 0-.537.213t-.213.537v.25q-1.05.25-1.65 1.15t-.6 2V15H8v1Zm-4 5V9l8-6l8 6v12H4Zm2-2h12v-9l-6-4.5L6 10v9Zm6-6.75Z"/>`),
		g.Group(children),
	)
}