package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Julia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<circle cx="29.1" cy="94.2" r="29.1" fill="currentColor"/><circle cx="98.9" cy="94.2" r="29.1" fill="currentColor"/><circle cx="64" cy="33.8" r="29.1" fill="currentColor"/>`),
		g.Group(children),
	)
}