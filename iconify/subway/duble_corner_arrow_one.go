package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DubleCornerArrowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m232.7.1l-93 93H349l-93 93l69.8 69.8l93-93v209.4l93-93L512 0L232.7.1zM256 325.8L186.2 256l-93 93V139.7l-93 93L0 512l279.3-.1l93-93H163l93-93.1z"/>`),
		g.Group(children),
	)
}