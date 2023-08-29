package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M467 0L117 331l374 404h-94L71 384v351H0V0h71v280L368 0h99z"/>`),
		g.Group(children),
	)
}