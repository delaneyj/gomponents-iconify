package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageAlertVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 9v11h-2v-9H2v9H0V9l10-4l10 4m-3 3H3v2h14v-2m0 3H3v2h14v-2m5 0v-5h2v5h-2m0 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}