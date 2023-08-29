package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBaseAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 8h5v2H7zm0 4h5v2H7zm0 4h5v2H7zm13-8h5v2h-5zm0 4h5v2h-5zm0 4h5v2h-5z"/><path fill="currentColor" d="M28 4H4a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2ZM4 6h11v22H4Zm13 22V6h11v22Z"/>`),
		g.Group(children),
	)
}