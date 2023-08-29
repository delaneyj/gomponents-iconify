package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandshakeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.1 21.4L10.7 20l5.325-5.325l-.725-.725l-5.325 5.325l-1.4-1.4L13.9 12.55l-.7-.7l-5.325 5.325l-1.4-1.4l5.3-5.35l-.725-.7l-5.325 5.325l-1.4-1.4l6.275-6.3l5.175 5.175L18.6 9.7l-5.9-5.9l2.05-2.05l8.5 8.5L12.1 21.4ZM3.525 13L.75 10.225L9.225 1.75l7.95 7.95l-1.425 1.425l-5.175-5.175L3.525 13Z"/>`),
		g.Group(children),
	)
}