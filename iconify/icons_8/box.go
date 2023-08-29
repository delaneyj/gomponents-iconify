package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 5v6h1v16h22V11h1V5H4zm2 2h20v2H6V7zm1 4h18v14H7V11zm5.813 2A1 1 0 0 0 13 15h6a1 1 0 1 0 0-2h-6a1 1 0 0 0-.094 0a1 1 0 0 0-.094 0z"/>`),
		g.Group(children),
	)
}