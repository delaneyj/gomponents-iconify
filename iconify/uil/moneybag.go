package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moneybag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7-9h-3V5a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Zm-9-1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4ZM4 9a1 1 0 0 1 1-1h1a2 2 0 0 1-2 2Zm1 11a1 1 0 0 1-1-1v-1a2 2 0 0 1 2 2Zm15-1a1 1 0 0 1-1 1h-1a2 2 0 0 1 2-2Zm0-3a4 4 0 0 0-4 4H8a4 4 0 0 0-4-4v-4a4 4 0 0 0 4-4h8a4 4 0 0 0 4 4Zm0-6a2 2 0 0 1-2-2h1a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}