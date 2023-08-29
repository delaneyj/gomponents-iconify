package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SneakerMoveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M256 200.41A15.91 15.91 0 0 1 240 216h-92.69a15.93 15.93 0 0 1-11.26-4.63L28.78 107.42l-.09-.09a16 16 0 0 1 0-22.62l64-64.12l.15-.14a15.91 15.91 0 0 1 22.35.27L123.4 29a16 16 0 0 1 4.66 10.54c1.13 22.83 16.91 38.26 41.19 40.26A16.13 16.13 0 0 1 184 95.7V108a4 4 0 0 1-4 4h-28a8 8 0 0 0-8 8.53a8.18 8.18 0 0 0 8.25 7.47h28a4 4 0 0 1 4 3.55a31.31 31.31 0 0 0 1.64 7.14a4 4 0 0 1-3.77 5.3H160a8 8 0 0 0-8 8.53a8.17 8.17 0 0 0 8.25 7.47H216a40 40 0 0 1 40 40.42ZM72 176a8 8 0 0 0-8-8H32a8 8 0 0 0 0 16h32a8 8 0 0 0 8-8Zm24 24H48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}