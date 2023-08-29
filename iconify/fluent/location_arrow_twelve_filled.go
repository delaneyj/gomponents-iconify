package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.64 1.069c.808-.311 1.603.484 1.292 1.292l-3.076 7.997c-.349.906-1.654.835-1.9-.104l-.803-3.05a.5.5 0 0 0-.357-.356l-3.05-.803c-.938-.247-1.01-1.552-.104-1.9L9.64 1.069Z"/>`),
		g.Group(children),
	)
}