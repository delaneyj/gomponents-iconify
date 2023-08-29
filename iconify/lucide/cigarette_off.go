package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CigaretteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M12 12H2v4h14m6-4v4m-4-4h-.5M7 12v4m11-8c0-2.5-2-2.5-2-5m6 5c0-2.5-2-2.5-2-5"/>`),
		g.Group(children),
	)
}