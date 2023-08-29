package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedTrianglePointedDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#F8312F" d="M15.134 22.5a1 1 0 0 0 1.732 0L22.928 12a1 1 0 0 0-.866-1.5H9.938a1 1 0 0 0-.866 1.5l6.062 10.5Z"/>`),
		g.Group(children),
	)
}