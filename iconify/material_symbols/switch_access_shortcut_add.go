package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchAccessShortcutAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 16v-2h-2v-2h2v-2h2v2h2v2h-2v2h-2Zm-3 6q-3.175-1.2-5.088-3.95T8 11.9q0-2.275.9-4.313T11.45 4H8V2h7v7h-2V5.3q-1.425 1.275-2.212 2.988T10 11.9q0 2.55 1.35 4.688T15 19.825V22Z"/>`),
		g.Group(children),
	)
}