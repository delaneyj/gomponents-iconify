package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NfcTap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 10a2 2 0 0 1 2 2a2 2 0 0 1-2 2a2 2 0 0 1-2-2a2 2 0 0 1 2-2M4 4h7a2 2 0 0 1 2 2v3h-2V6H4v5h2V9l3 3l-3 3v-2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2m16 16h-7a2 2 0 0 1-2-2v-3h2v3h7v-5h-2v2l-3-3l3-3v2h2a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}