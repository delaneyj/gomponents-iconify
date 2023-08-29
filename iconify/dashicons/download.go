package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.01 4v6h2V2H4v8h2.01V4h8zm-2 2v6h3l-5 6l-5-6h3V6h4z"/>`),
		g.Group(children),
	)
}