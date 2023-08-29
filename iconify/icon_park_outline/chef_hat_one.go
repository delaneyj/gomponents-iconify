package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChefHatOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 24.125a8.64 8.64 0 1 1 3.857-16.837A8.625 8.625 0 0 1 23.64 4a8.627 8.627 0 0 1 6.919 3.464A8.64 8.64 0 1 1 35 24.124V40a2 2 0 0 1-2 2H15a2 2 0 0 1-2-2V24.125ZM13 31h22m-15-6v6m15-3v6m-22-6v6"/>`),
		g.Group(children),
	)
}