package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodFrownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm1.82 5.035a7.278 7.278 0 0 1 4.368-1.092l.498.035l.07-.998l-.5-.034a8.278 8.278 0 0 0-4.966 1.241l-.424.266l.531.847l.424-.265ZM11 6h-1V5h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}