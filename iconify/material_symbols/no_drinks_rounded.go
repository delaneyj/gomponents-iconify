package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoDrinksRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21q-.425 0-.713-.288T6 20q0-.425.288-.713T7 19h4v-5l-1.2-1.35l-7.7-7.725q-.275-.275-.275-.687t.275-.713q.3-.3.713-.3t.712.3L20.5 20.5q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3L13 15.85V19h4.025q.425 0 .7.288T18 20q0 .425-.288.713T17 21H7ZM5.85 3h13.7q.6 0 1.025.438T21 4.474q0 .275-.088.5t-.262.425l-5.85 6.55L9.85 7h6.7l1.8-2H7.85l-2-2Z"/>`),
		g.Group(children),
	)
}