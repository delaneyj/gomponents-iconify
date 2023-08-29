package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 1h4.764l3 11h10.515l3.088-9.265l1.898.633L19.72 14H7.78l-.5 2H22v2H4.72l1.246-4.989L3.236 3H0V1Zm14 1v3h3v2h-3v3h-2V7H9V5h3V2h2ZM4 21a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm14 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/>`),
		g.Group(children),
	)
}