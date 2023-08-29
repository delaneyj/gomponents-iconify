package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.807 5.687c.298.245.546.55.727.897a.236.236 0 0 1-.091.307l-6.266 3.88a4.25 4.25 0 0 1-4.4.045L3.47 7.088a.236.236 0 0 1-.103-.293a2.89 2.89 0 0 1 .823-1.106v-.003l.012-.007a2.9 2.9 0 0 1 .894-.496l4.11-2.284a5.75 5.75 0 0 1 5.585 0l4.105 2.28c.334.114.641.286.908.505l.003.003Z"/><path fill="currentColor" d="M2.989 8.954a.248.248 0 0 1 .373-.187l5.653 3.34a5.75 5.75 0 0 0 5.951-.061l5.645-3.495a.248.248 0 0 1 .377.183a30.35 30.35 0 0 1-.161 7.78a2.888 2.888 0 0 1-2.606 2.447l-1.51.131a54.394 54.394 0 0 1-9.422 0l-1.51-.131a2.888 2.888 0 0 1-2.606-2.448a30.352 30.352 0 0 1-.184-7.559Z"/>`),
		g.Group(children),
	)
}