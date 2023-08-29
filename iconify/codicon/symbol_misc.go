package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SymbolMisc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 2h8v4c.341.035.677.112 1 .23V1H3v8.48l1-1.75V2zm2.14 8L5 8L4 9.75L3.29 11L1 15h8l-2.29-4l-.57-1zm-3.42 4l1.72-3L5 10l.56 1l1.72 3H2.72zm6.836-6.41a3.5 3.5 0 1 1 3.888 5.82a3.5 3.5 0 0 1-3.888-5.82zm.555 4.989a2.5 2.5 0 1 0 2.778-4.157a2.5 2.5 0 0 0-2.778 4.157z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}