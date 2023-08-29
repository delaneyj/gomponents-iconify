package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageOpenVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 9v11h-2v-9H4v9H2V9l10-4l10 4m-3 3H5v2h14v-2Z"/>`),
		g.Group(children),
	)
}