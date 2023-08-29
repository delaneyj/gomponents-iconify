package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0S3 8.2 3 11s2.2 5 5 5s5-2.2 5-5S8 0 8 0zm.9 14.9l-.2-1c1.4-.3 2.4-1.7 2.4-3.2c0-.3-.1-1.1-.8-2.6l.9-.4c.6 1.4.8 2.4.8 3c0 2-1.3 3.8-3.1 4.2z"/>`),
		g.Group(children),
	)
}