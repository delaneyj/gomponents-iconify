package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 28a1 1 0 0 1-.5-.13l-19-11a1 1 0 0 1 0-1.74l19-11a1 1 0 0 1 1 0A1 1 0 0 1 28 5v22a1 1 0 0 1-1 1zM10 16l16 9.27V6.73zM2 4h2v24H2z"/>`),
		g.Group(children),
	)
}