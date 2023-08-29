package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2v5.938l-1.142.542a5 5 0 1 0 4.284 0L8 7.938V2H6zM4 6.674V0h6v6.674a7 7 0 1 1-6 0zM3.535 11h6.93a4 4 0 1 1-6.93 0zM4 0h6a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2z"/>`),
		g.Group(children),
	)
}