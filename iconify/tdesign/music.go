package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Music(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2.913V18a3 3 0 1 1-2-2.83v-5.083L8 10.92V19a3 3 0 1 1-2-2.83V4.08l14-1.167ZM6 19a1 1 0 1 0-2 0a1 1 0 0 0 2 0ZM8 8.913l10-.833V5.087L8 5.92v2.993ZM18 18a1 1 0 1 0-2 0a1 1 0 0 0 2 0Z"/>`),
		g.Group(children),
	)
}