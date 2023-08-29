package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackBelowSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 14a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1h11ZM5 10H2.5a.5.5 0 0 1-.5-.5V6h3v4Zm1 0h4V6H6v4Zm8-4v3.5a.5.5 0 0 1-.5.5H11V6h3Zm-3-1h3v-.5A2.5 2.5 0 0 0 11.5 2H11v3Zm-1-3H6v3h4V2ZM4.5 2H5v3H2v-.5A2.5 2.5 0 0 1 4.5 2Z"/>`),
		g.Group(children),
	)
}