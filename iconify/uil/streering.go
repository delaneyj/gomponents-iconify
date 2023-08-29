package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Streering(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm9.71-2.36a10 10 0 0 0-19.4 0a9.75 9.75 0 0 0 0 4.72a10 10 0 0 0 7.3 7.34a9.67 9.67 0 0 0 4.7 0a10 10 0 0 0 7.31-7.31a9.75 9.75 0 0 0 0-4.72ZM12 4a8 8 0 0 1 7.41 5H4.59A8 8 0 0 1 12 4Zm-8 8a8.26 8.26 0 0 1 .07-1H6v2H4.07A8.26 8.26 0 0 1 4 12Zm5 7.41A8 8 0 0 1 4.59 15H7a2 2 0 0 1 2 2Zm4 .52a8.26 8.26 0 0 1-1 .07a8.26 8.26 0 0 1-1-.07V18h2Zm.14-3.93h-2.28A4 4 0 0 0 8 13.14V11h8v2.14A4 4 0 0 0 13.14 16ZM15 19.41V17a2 2 0 0 1 2-2h2.41A8 8 0 0 1 15 19.41ZM19.93 13H18v-2h1.93a8.26 8.26 0 0 1 .07 1a8.26 8.26 0 0 1-.07 1Z"/>`),
		g.Group(children),
	)
}