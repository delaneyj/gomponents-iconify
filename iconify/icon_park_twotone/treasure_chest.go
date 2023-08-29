package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreasureChest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTreasureChest0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M42 4H6a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path stroke-linecap="round" d="M4 24h13m14 0h13"/><path fill="#555" d="M24 31a7 7 0 1 0 0-14a7 7 0 0 0 0 14Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTreasureChest0)"/>`),
		g.Group(children),
	)
}