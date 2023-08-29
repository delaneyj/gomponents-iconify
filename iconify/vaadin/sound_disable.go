package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundDisable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 5H0v6h4l5 4V1zm11.9.6l-.8-.7l-2.3 2.4l-2.4-2.4l-.8.7L12 8l-2.4 2.4l.8.7l2.4-2.4l2.3 2.4l.8-.7L13.5 8z"/>`),
		g.Group(children),
	)
}