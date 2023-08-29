package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1V2.5Zm2 2v4.25h18V7H10.52l-2-2.5H3Zm18 6.25H3V19h18v-8.25Z"/>`),
		g.Group(children),
	)
}