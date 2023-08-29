package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vinyl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.999.8A9.2 9.2 0 0 0 .8 10.001a9.2 9.2 0 0 0 18.399 0A9.2 9.2 0 0 0 9.999.8zM10 13.001a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/>`),
		g.Group(children),
	)
}