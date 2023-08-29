package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ninja(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.44 5.34l-.06-.07a10 10 0 0 0-14.76 0l-.06.07A10 10 0 1 0 22 12a9.93 9.93 0 0 0-2.56-6.66ZM12 4a7.87 7.87 0 0 1 3.86 1H8.14A7.87 7.87 0 0 1 12 4ZM5.76 7h12.48a8 8 0 0 1 1.69 4H4.07a8 8 0 0 1 1.69-4ZM12 20a8 8 0 0 1-7.93-7h15.86A8 8 0 0 1 12 20ZM8 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm6 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm-3 9.93a1 1 0 0 0 .49.13a1 1 0 0 0 .87-.51A3 3 0 0 1 15 16a1 1 0 0 0 0-2a5 5 0 0 0-4.37 2.57a1 1 0 0 0 .37 1.36Z"/>`),
		g.Group(children),
	)
}