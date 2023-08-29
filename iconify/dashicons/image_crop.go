package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageCrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 12v3h-4v4h-3v-4H4V7H0V4h4V0h3v4h7l3-3l1 1l-3 3v7h4zm-8-5H7v4zm-3 5h4V8z"/>`),
		g.Group(children),
	)
}