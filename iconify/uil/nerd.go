package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nerd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.41 8.65v-.06a10 10 0 0 0-18.78-.06a.85.85 0 0 0-.08.24A9.87 9.87 0 0 0 2 12a10 10 0 1 0 19.41-3.35ZM12 4a8 8 0 0 1 6.92 4h-1.2a3 3 0 0 0-4.62.22A3.17 3.17 0 0 0 12 8a3.17 3.17 0 0 0-1.1.22A3 3 0 0 0 6.28 8h-1.2A8 8 0 0 1 12 4Zm4.5 6a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm-7 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1ZM12 20a8 8 0 0 1-8-8a8.24 8.24 0 0 1 .26-2H5.5a3 3 0 0 0 6 .18a1 1 0 0 1 1 0a3 3 0 0 0 6-.18h1.24a8.24 8.24 0 0 1 .26 2a8 8 0 0 1-8 8Zm2.36-5.77a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54Z"/>`),
		g.Group(children),
	)
}