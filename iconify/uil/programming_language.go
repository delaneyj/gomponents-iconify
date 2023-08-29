package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProgrammingLanguage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 3l-3 15.1l-9.1 3l-7.9-3l.8-4.1h3.4l-.4 1.7l4.8 1.8l5.5-1.8l.8-3.8H3.2l.7-3.4h13.6l.5-2.1H4.3L5 3h17z"/>`),
		g.Group(children),
	)
}