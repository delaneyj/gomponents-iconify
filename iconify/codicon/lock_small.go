package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m3 8l1-1h8l1 1v5l-1 1H4l-1-1V8Zm1 0v5h8V8H4Zm7-1V5a3 3 0 0 0-6 0v2h1V5a2 2 0 1 1 4 0v2h1Z"/>`),
		g.Group(children),
	)
}