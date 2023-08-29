package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PausePast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 11h-2v11h2V11zm6 0h-2v11h2V11z"/><path fill="currentColor" d="M16 2A13.916 13.916 0 0 0 6 6.234V2H4v8h8V8H7.078A11.982 11.982 0 1 1 4 16H2A14 14 0 1 0 16 2Z"/>`),
		g.Group(children),
	)
}