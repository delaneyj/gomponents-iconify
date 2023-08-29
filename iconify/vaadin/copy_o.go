package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopyO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13 3h-3L7 0H0v13h6v3h10V6l-3-3zM7 1l2 2H7V1zM1 12V1h5v3h3v8H1zm14 3H7v-2h3V4h2v3h3v8zm-2-9V4l2 2h-2z"/>`),
		g.Group(children),
	)
}