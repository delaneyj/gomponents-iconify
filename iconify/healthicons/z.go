package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Z(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M29.73 14H14a2 2 0 0 1 0-4h20a2 2 0 0 1 1.536 3.28L18.27 34H34a2 2 0 1 1 0 4H14a2 2 0 0 1-1.536-3.28L29.73 14Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}