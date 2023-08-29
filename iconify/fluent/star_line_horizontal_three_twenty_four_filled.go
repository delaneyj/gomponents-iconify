package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarLineHorizontalThreeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.788 3.102c.495-1.003 1.926-1.003 2.421 0l1.687 3.425a2.25 2.25 0 0 0-.678 4.223a2.25 2.25 0 0 0 0 4a2.25 2.25 0 0 0-.767 3.351l-6.168 3.241c-.99.52-2.148-.32-1.96-1.423l.901-5.251l-3.815-3.72c-.801-.78-.359-2.141.748-2.302L8.43 7.88l2.358-4.778ZM14.5 16.75a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.747.747 0 0 1-.75-.75Zm0-4a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1-.75-.75Zm0-4a.75.75 0 0 1 .75-.75h6a.75.75 0 0 1 0 1.5h-6a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}