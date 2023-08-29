package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M414 354q-18-18-41-11l-32-32q43-53 43-119q0-80-56-136T192 0T56 56T0 192t56 136t136 56q70 0 119-43l32 32q-6 24 11 41l85 85q13 13 30 13q18 0 30-13q13-13 13-30t-13-30zm-222-13q-62 0-105.5-43.5T43 192T86.5 86.5T192 43t105.5 43.5T341 192t-43.5 105.5T192 341zm64-170h-43v-43q0-21-21-21t-21 21v43h-43q-21 0-21 21t21 21h43v43q0 21 21 21t21-21v-43h43q21 0 21-21t-21-21z"/>`),
		g.Group(children),
	)
}