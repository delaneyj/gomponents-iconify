package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V15.437l1.72-1.718l.28-.314V3H6zm2 2h14v8.406l.28.313L24 15.436V27H8V5zm16 0h2v7.563l-1 1l-1-1V5z"/>`),
		g.Group(children),
	)
}