package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NatureEcologyVolcanoEruptionEruptMountainVolcanoLavaMagmaExplosion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 13.5s-4-2-4-5h-5c0 3-4 5-4 5m10-8a2 2 0 0 0 0-4a2 2 0 0 0-1.24.43a2.5 2.5 0 0 0-4.52 0A2 2 0 0 0 3.5 1.5a2 2 0 0 0 0 4m2-1v4m3-4v4"/>`),
		g.Group(children),
	)
}