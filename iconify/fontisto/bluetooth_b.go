package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.446 19.058l2.32-2.303l-2.32-2.303zm0-9.508l2.32-2.303l-2.32-2.303zM7.875 12l4.768 4.768L5.424 24v-9.52l-3.978 3.964L0 16.995L4.982 12L0 7.005l1.446-1.446l3.978 3.962V.001l7.216 7.232z"/>`),
		g.Group(children),
	)
}