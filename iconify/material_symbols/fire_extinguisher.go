package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireExtinguisher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22q-.825 0-1.413-.588T7 20v-1h10v1q0 .825-.588 1.413T15 22H9Zm-2-4v-5h10v5H7Zm.1-6q.25-1.15 1.05-2.163t2.025-1.487q-.25-.2-.45-.438t-.35-.512L5 6.5v-1l4.375-.9q.375-.725 1.063-1.163T12 3q.525 0 1 .175t.85.475L17 3v6l-3.175-.65q1.175.475 1.988 1.438T16.9 12H7.1ZM12 7q.425 0 .713-.288T13 6q-.025-.45-.3-.725T12 5q-.425 0-.713.288T11 6q0 .425.288.713T12 7Z"/>`),
		g.Group(children),
	)
}