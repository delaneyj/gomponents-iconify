package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kinopoisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.51 21.51 0 0 0 24 2.5Zm-4.77 5.85a5.32 5.32 0 1 1-5.31 5.31a5.32 5.32 0 0 1 5.31-5.31Zm15 5.52a5.31 5.31 0 1 1-5.32 5.31a5.31 5.31 0 0 1 5.35-5.31ZM24 21.33A2.67 2.67 0 0 1 26.69 24h0A2.68 2.68 0 0 1 24 26.67h0a2.67 2.67 0 1 1 0-5.34Zm-10.26 2.12a5.32 5.32 0 1 1-5.32 5.32h0a5.32 5.32 0 0 1 5.32-5.32ZM28.79 29a5.31 5.31 0 1 1-5.31 5.31A5.32 5.32 0 0 1 28.79 29Z"/>`),
		g.Group(children),
	)
}