package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaSpreadsheet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12 2l4 4v12H4V2h8zm-1 4V3H5v3h6zM8 8V7H5v1h3zm3 0V7H9v1h2zm4 0V7h-3v1h3zm-7 2V9H5v1h3zm3 0V9H9v1h2zm4 0V9h-3v1h3zm-7 2v-1H5v1h3zm3 0v-1H9v1h2zm4 0v-1h-3v1h3zm-7 2v-1H5v1h3zm3 0v-1H9v1h2zm4 0v-1h-3v1h3zm-7 2v-1H5v1h3zm3 0v-1H9v1h2z"/>`),
		g.Group(children),
	)
}