package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M221 88q-26 0-51 9l-33-32q40-20 84-20q79 0 135.5 56.5T413 237q0 44-20 84l-32-32q9-26 9-52q0-62-43.5-105.5T221 88zm213-6l-27 33l-99-83l28-32zM27 9l21 21l372 372l-27 27l-47-47q-54 47-125 47q-80 0-136-56T29 237q0-71 47-125L59 95l-24 20L5 84l23-19L0 36zm289 343L106 142q-35 42-35 95q0 62 44 106t106 44q54 0 95-35zM136 30l-18 15l-31-30l19-15z"/>`),
		g.Group(children),
	)
}