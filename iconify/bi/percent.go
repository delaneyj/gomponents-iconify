package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M13.442 2.558a.625.625 0 0 1 0 .884l-10 10a.625.625 0 1 1-.884-.884l10-10a.625.625 0 0 1 .884 0zM4.5 6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5zm7 6a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5z"/>`),
		g.Group(children),
	)
}