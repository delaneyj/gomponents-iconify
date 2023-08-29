package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAnnotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M13.71 12.29L7.42 6H14V4H4v9.99l2 .02v-6.6l6.29 6.29l1.42-1.41z" fill="currentColor"/><path d="M28 10H18v2h10v16H12V18h-2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2z" fill="currentColor"/><path d="M19 25h2v-7h3v-2h-8v2h3v7z" fill="currentColor"/>`),
		g.Group(children),
	)
}