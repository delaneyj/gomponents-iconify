package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feedback(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 2h16c.55 0 1 .45 1 1v14c0 .55-.45 1-1 1H2c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1zm15 14V7H3v9h14zM4 8v1h3V8H4zm4 0v3h8V8H8zm-4 4v1h3v-1H4zm4 0v3h8v-3H8z"/>`),
		g.Group(children),
	)
}