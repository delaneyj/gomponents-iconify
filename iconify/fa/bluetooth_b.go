package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m564 1423l173-172l-173-172v344zm0-710l173-172l-173-172v344zm32 183l356 356l-539 540v-711l-297 296L8 1269l372-373L8 523l108-108l297 296V0l539 540z"/>`),
		g.Group(children),
	)
}