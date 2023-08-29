package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VivoThemes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.073 4.5h21.854A2.567 2.567 0 0 1 37.5 7.073l-.464 14.917H10.783L10.5 7.073C10.476 5.648 11.648 4.5 13.073 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.036 21.99v4.632c.172 1.885-2.133 1.862-3.628 1.87c-5.494.028-6.634 2.567-6.469 5.291c.155 2.572.197 2.898.232 4.349c.085 3.487-1.595 5.368-3.262 5.368s-3.346-1.88-3.261-5.368c.035-1.451.076-1.777.232-4.349c.164-2.724-.976-5.263-6.47-5.291c-1.495-.008-3.8.015-3.627-1.87V21.99"/>`),
		g.Group(children),
	)
}