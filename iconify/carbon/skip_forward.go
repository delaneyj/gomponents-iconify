package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 28a1 1 0 0 1-1-1V5a1 1 0 0 1 .5-.87a1 1 0 0 1 1 0l19 11a1 1 0 0 1 0 1.74l-19 11A1 1 0 0 1 5 28zM6 6.73v18.54L22 16zM28 4h2v24h-2z"/>`),
		g.Group(children),
	)
}