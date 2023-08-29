package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeenhereRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22.5q-.325 0-.625-.1t-.575-.3l-6-4.5q-.375-.275-.587-.7T4 16V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v12q0 .475-.213.9t-.587.7l-6 4.5q-.275.2-.575.3t-.625.1Zm-1.05-10.35l-1.4-1.4q-.3-.3-.7-.288t-.7.288q-.3.3-.313.713t.288.712L10.25 14.3q.3.3.7.3t.7-.3l4.25-4.25q.3-.3.287-.7t-.287-.7q-.3-.3-.713-.313t-.712.288L10.95 12.15Z"/>`),
		g.Group(children),
	)
}