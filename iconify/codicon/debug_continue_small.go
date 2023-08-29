package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugContinueSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 2H3v12h1V2Zm3.29.593L6.5 3v10l.79.407l7-5v-.814l-7-5ZM13.14 8L7.5 12.028V3.972L13.14 8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}