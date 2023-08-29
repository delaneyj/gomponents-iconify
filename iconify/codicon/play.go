package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.78 2L3 2.41v12l.78.42l9-6V8l-9-6zM4 13.48V3.35l7.6 5.07L4 13.48z"/>`),
		g.Group(children),
	)
}