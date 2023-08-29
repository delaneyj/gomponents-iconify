package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.667 9.417a2.583 2.583 0 1 0 0 5.166h2.583v-2.531a4.511 4.511 0 0 1 0-.104V9.417H8.667Zm2.583-1.5H8.667a2.583 2.583 0 0 1 0-5.167h2.583v5.167Zm1.5-5.167v5.167h2.583a2.584 2.584 0 0 0 0-5.167H12.75Zm2.583 6.666a2.583 2.583 0 0 0-2.583 2.542v.083a2.583 2.583 0 1 0 2.583-2.625Zm-6.666 6.667a2.584 2.584 0 1 0 2.583 2.584v-2.584H8.667Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}