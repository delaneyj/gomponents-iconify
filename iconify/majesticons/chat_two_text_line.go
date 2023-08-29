package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTwoTextLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8h8m-8 3h8m-8 3h3m11-3c0 4.418-4.477 8-10 8c-1.3 0-2.54-.198-3.68-.559C6.528 20.021 5 21 2 21c1-1 2.27-2.35 2.801-4.447C3.067 15.114 2 13.157 2 11c0-4.418 4.477-8 10-8s10 3.582 10 8z"/>`),
		g.Group(children),
	)
}