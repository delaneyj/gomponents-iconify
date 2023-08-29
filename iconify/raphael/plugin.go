package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plugin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m26.33 15.836l-3.893-1.545l3.136-7.9a1.374 1.374 0 1 0-2.556-1.014l-3.136 7.9L15 11.34l3.136-7.9a1.375 1.375 0 0 0-2.555-1.016l-3.135 7.9l-3.89-1.543l-1.615 4.067l2.15.854l-2.537 6.392a3.002 3.002 0 0 0 1.683 3.895l1.626.646l-.877 2.207a2 2 0 0 0 1.122 2.596l.93.37a2.001 2.001 0 0 0 2.596-1.122l.877-2.207l1.858.737a3.002 3.002 0 0 0 3.896-1.682l2.535-6.39l1.917.76l1.613-4.066z"/>`),
		g.Group(children),
	)
}