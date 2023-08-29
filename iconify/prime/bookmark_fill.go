package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.75 3.25h-7.5A2.75 2.75 0 0 0 5.5 6v14a.75.75 0 0 0 1.18.62L12 16.91l5.32 3.71a.75.75 0 0 0 .43.13a.85.85 0 0 0 .35-.08a.77.77 0 0 0 .4-.67V6a2.75 2.75 0 0 0-2.75-2.75Z"/>`),
		g.Group(children),
	)
}