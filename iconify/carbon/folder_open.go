package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 8h-7.172l-3.414-3.414A2 2 0 0 0 16 4H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2ZM8 26V14h8v6.17l-2.59-2.58L12 19l5 5l5-5l-1.41-1.41L18 20.17V14a2.002 2.002 0 0 0-2-2H8a2.002 2.002 0 0 0-2 2v12H4V6h12l4 4h8v2h-6v2h6v12Z"/>`),
		g.Group(children),
	)
}