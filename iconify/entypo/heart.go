package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.19 4.155c-1.672-1.534-4.383-1.534-6.055 0L10 5.197L8.864 4.155c-1.672-1.534-4.382-1.534-6.054 0c-1.881 1.727-1.881 4.52 0 6.246L10 17l7.19-6.599c1.88-1.726 1.88-4.52 0-6.246z"/>`),
		g.Group(children),
	)
}