package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DubleCornerArrowBlodTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m511.9 232.7l-93-93V349l-93-93l-69.9 69.8l93 93H139.7l93 93l279.3.2l-.1-279.3zM256 186.2l-93-93h209.4l-93-93L0 0l.1 279.3l93 93V163l93 93l69.9-69.8z"/>`),
		g.Group(children),
	)
}