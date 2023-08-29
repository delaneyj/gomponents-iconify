package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleSheets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.4 4.5a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h27.2a2 2 0 0 0 2-2v-27h-8a2 2 0 0 1-2-2v-8ZM16.55 18h14.9a1 1 0 0 1 1 1v19.7a1 1 0 0 1-1 1h-14.9a1 1 0 0 1-1-1V19a1 1 0 0 1 1-1ZM29.61 4.5l9.99 9.99m-21.13 8.83h11.06m-11.06 5.47h11.06m-11.06 5.49h11.06"/>`),
		g.Group(children),
	)
}