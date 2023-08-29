package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChooseItem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 6h2v20h-2zM17 6l-1.43 1.393L23.15 15H2v2h21.15l-7.58 7.573L17 26l10-10L17 6z"/>`),
		g.Group(children),
	)
}