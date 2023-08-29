package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TicketLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M226.79 102.52A14 14 0 0 0 238 88.81V64a14 14 0 0 0-14-14H32a14 14 0 0 0-14 14v24.81a14 14 0 0 0 11.21 13.71a26 26 0 0 1 0 51A14 14 0 0 0 18 167.19V192a14 14 0 0 0 14 14h192a14 14 0 0 0 14-14v-24.81a14 14 0 0 0-11.21-13.71a26 26 0 0 1 0-51ZM30 192v-24.81a2 2 0 0 1 1.59-1.95a38 38 0 0 0 0-74.48a2 2 0 0 1-1.59-2V64a2 2 0 0 1 2-2h58v132H32a2 2 0 0 1-2-2Zm194.41-26.76a2 2 0 0 1 1.59 1.95V192a2 2 0 0 1-2 2H102V62h122a2 2 0 0 1 2 2v24.81a2 2 0 0 1-1.59 2a38 38 0 0 0 0 74.48Z"/>`),
		g.Group(children),
	)
}