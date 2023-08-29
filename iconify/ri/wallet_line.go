package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.005 7h3a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h15v4Zm-14 2v10h16V9h-16Zm0-4v2h12V5h-12Zm11 8h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}