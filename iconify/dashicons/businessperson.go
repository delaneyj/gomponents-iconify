package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Businessperson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.2 10L11 13l-1-1.4L9 13l-2.2-3C3 11 3 13 3 16.9c0 0 3 1.1 6.4 1.1h1.2c3.4-.1 6.4-1.1 6.4-1.1c0-3.9 0-5.9-3.8-6.9zm-3.2.7L8.4 10l1.6 1.6l1.6-1.6l-1.6.7zm0-8.6c-1.9 0-3 1.8-2.7 3.8c.3 2 1.3 3.4 2.7 3.4s2.4-1.4 2.7-3.4c.3-2.1-.8-3.8-2.7-3.8z"/>`),
		g.Group(children),
	)
}