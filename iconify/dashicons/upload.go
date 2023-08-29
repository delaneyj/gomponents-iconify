package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 14V8H5l5-6l5 6h-3v6H8zm-2 2v-6H4v8h12.01v-8H14v6H6z"/>`),
		g.Group(children),
	)
}