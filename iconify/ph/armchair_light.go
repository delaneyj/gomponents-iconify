package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArmchairLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214 90.48V72a38 38 0 0 0-38-38H80a38 38 0 0 0-38 38v18.48a38 38 0 0 0 0 75.05V200a14 14 0 0 0 14 14h144a14 14 0 0 0 14-14v-34.47a38 38 0 0 0 0-75ZM80 46h96a26 26 0 0 1 26 26v18.48A38.05 38.05 0 0 0 170 128v2H86v-2a38.05 38.05 0 0 0-32-37.52V72a26 26 0 0 1 26-26Zm128.35 108H208a6 6 0 0 0-6 6v40a2 2 0 0 1-2 2H56a2 2 0 0 1-2-2v-40a6 6 0 0 0-6-6h-.35A26 26 0 1 1 74 128v40a6 6 0 0 0 12 0v-26h84v26a6 6 0 0 0 12 0v-40a26 26 0 1 1 26.35 26Z"/>`),
		g.Group(children),
	)
}