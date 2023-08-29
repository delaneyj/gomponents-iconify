package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V15.437l1.719-1.718l.281-.313V3zm2 2h14v8.406l.281.313L24 15.438V27H8zm16 0h2v7.563l-1 1l-1-1z"/>`),
		g.Group(children),
	)
}