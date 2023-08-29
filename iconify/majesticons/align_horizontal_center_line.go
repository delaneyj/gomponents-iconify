package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalCenterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6H6a2 2 0 1 0 0 4h6m0-4h6a2 2 0 1 1 0 4h-6m0-4V3m0 7v4m0 0H9a2 2 0 1 0 0 4h3m0-4h3a2 2 0 1 1 0 4h-3m0 0v3"/>`),
		g.Group(children),
	)
}