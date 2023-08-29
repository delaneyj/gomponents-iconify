package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdReorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M96 348h320v36H96z" fill="currentColor"/><path d="M96 128h320v36H96z" fill="currentColor"/><path d="M96 200.7h320v35.6H96z" fill="currentColor"/><path d="M96 275.8h320v35.6H96z" fill="currentColor"/>`),
		g.Group(children),
	)
}