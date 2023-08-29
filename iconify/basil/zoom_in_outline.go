package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomInOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.2 6.95a.75.75 0 0 1 .75.75v1.75h1.75a.75.75 0 0 1 0 1.5h-1.75v1.75a.75.75 0 0 1-1.5 0v-1.75H7.7a.75.75 0 0 1 0-1.5h1.75V7.7a.75.75 0 0 1 .75-.75Z"/><path fill="currentColor" fill-rule="evenodd" d="M5.399 14.945a6.75 6.75 0 0 0 8.986.5l5.156 5.156a.75.75 0 1 0 1.06-1.06l-5.155-5.156a6.75 6.75 0 1 0-10.047.56Zm1.06-8.486a5.25 5.25 0 0 0 7.42 7.43l.005-.005l.005-.005a5.25 5.25 0 0 0-7.43-7.42Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}